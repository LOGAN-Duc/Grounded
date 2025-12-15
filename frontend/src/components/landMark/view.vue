<template>
  <div class="container mt-4">
    <NavMenu />

    <div class="d-flex justify-content-between align-items-center mb-3">
      <h2>Danh sách Landmarks</h2>
      <div>
        <button class="btn btn-primary" @click="goToCreate">+ Tạo Landmark</button>
      </div>
      <input type="text" class="form-control w-50" placeholder="Tìm theo tên..." v-model="search"/>
    </div>

    <table class="table table-bordered align-middle" v-if="!loading && landmarks.length">
      <thead>
        <tr>
          <th>Tên</th>
          <th>Mô tả</th>
          <th>Ảnh</th>
          <th>Hành động</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="lm in paginatedLandmarks" :key="lm.id">
          <!-- Tên Landmark bấm vào sẽ xem chi tiết -->
          <td>
            <span class="link-primary pointer" @click="viewLandmark(lm.id)">
              {{ lm.name }}
            </span>
          </td>
          <td>{{ lm.description || '-' }}</td>
          <td>
            <img v-if="lm.urlImage" :src="lm.urlImage" alt="" style="width:80px;height:80px;object-fit:cover;">
            <span v-else>Không có ảnh</span>
          </td>
          <td>
            <button class="btn btn-sm btn-warning" @click.stop="editLandmark(lm.id)">Edit</button>
            <button class="btn btn-sm btn-success ms-1" @click.stop="goToAddItemsResources(lm.id)">Thêm Item/Resource</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-else-if="!loading" class="text-center">Không có dữ liệu</div>
    <div v-else class="text-center">Đang tải dữ liệu...</div>

    <!-- Pagination -->
    <nav v-if="totalPages > 1" class="mt-3">
      <ul class="pagination justify-content-center">
        <li class="page-item" :class="{disabled: currentPage === 1}">
          <button class="page-link" @click="changePage(currentPage - 1)">Previous</button>
        </li>
        <li class="page-item" v-for="p in totalPages" :key="p" :class="{active: currentPage === p}">
          <button class="page-link" @click="changePage(p)">{{ p }}</button>
        </li>
        <li class="page-item" :class="{disabled: currentPage === totalPages}">
          <button class="page-link" @click="changePage(currentPage + 1)">Next</button>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
import axios from 'axios';
import NavMenu from '../layout/NavMenu.vue';

export default {
  name: 'LandmarkList',
  components: { NavMenu },
  data() {
    return {
      landmarks: [],
      search: '',
      loading: false,
      currentPage: 1,
      itemsPerPage: 10,
    };
  },
  computed: {
    filteredLandmarks() {
      return this.landmarks.filter(lm => lm.name.toLowerCase().includes(this.search.toLowerCase()));
    },
    paginatedLandmarks() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      return this.filteredLandmarks.slice(start, start + this.itemsPerPage);
    },
    totalPages() {
      return Math.ceil(this.filteredLandmarks.length / this.itemsPerPage);
    },
  },
  methods: {
    async fetchLandmarks() {
      this.loading = true;
      try {
        const res = await axios.get('http://localhost:9999/landmarks/');
        this.landmarks = (res.data.result || []).map(lm => {
          if (lm.urlImage) {
            const fileName = lm.urlImage.split('/').pop();
            lm.urlImage = '/items/' + fileName;
          }
          return lm;
        });
      } catch(err) {
        console.error('Fetch error:', err);
      } finally {
        this.loading = false;
      }
    },
    goToCreate() { this.$router.push('/landmarks/create/'); },
    editLandmark(id) { this.$router.push(`/landmarks/edit/${id}`); },
    goToAddItemsResources(landmarkId) {
      this.$router.push(`/landmarks/${landmarkId}/add-items-resources`);
    },
    viewLandmark(id) {
      this.$router.push(`/landmarks/view/${id}`);
    },
    changePage(page) {
      if (page < 1 || page > this.totalPages) return;
      this.currentPage = page;
    }
  },
  watch: {
    search() {
      this.currentPage = 1;
    }
  },
  mounted() {
    this.fetchLandmarks();
  }
};
</script>

<style scoped>
.table td {
  vertical-align: middle;
}
.list-group-item:hover {
  background-color: #f1f1f1;
}
.pointer {
  cursor: pointer;
}
.link-primary {
  color: #0d6efd;
  text-decoration: underline;
}
</style>
