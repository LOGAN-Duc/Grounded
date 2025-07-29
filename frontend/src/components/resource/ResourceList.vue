<template>
  <div class="container mt-4">
    <NavMenu />
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h2 class="mb-0">Danh sách Resources</h2>
      <div class="d-flex gap-2 w-50">
        <input
          type="text"
          class="form-control"
          placeholder="Tìm theo tên..."
          v-model="search"
        />
      </div>
    </div>

    <table class="table table-bordered align-middle">
      <thead class="thead-light">
        <tr>
          <th style="white-space: nowrap; width: 200px;">Tên</th>
          <th style="width: 100px;">Mã</th>
          <th style="width: 200px;">Loại Resource</th>
          <th style="width: 150px;">Hành động</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="resource in paginatedResources" :key="resource.id">
          <td style="white-space: nowrap;">{{ resource.name }}</td>
          <td :title="resource.code">{{ resource.code.substring(0, 10) }}...</td>
          <td>{{ resource.resourceType.name || 'Không xác định' }}</td>
          <td>
            <button class="btn btn-sm btn-warning" @click="editResource(resource.id)">Edit</button>
          </td>
        </tr>
        <tr v-if="paginatedResources.length === 0">
          <td colspan="4" class="text-center">Không tìm thấy resource nào</td>
        </tr>
      </tbody>
    </table>

    <div class="d-flex justify-content-between align-items-center mt-3">
      <button class="btn btn-secondary" @click="prevPage" :disabled="currentPage === 1">Trước</button>
      <span>Trang {{ currentPage }} / {{ totalPages }}</span>
      <button class="btn btn-secondary" @click="nextPage" :disabled="currentPage === totalPages">Sau</button>
    </div>

    <div class="mt-3">
      <button class="btn btn-primary" @click="goToCreate">+ Tạo Resource</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import NavMenu from '../layout/NavMenu.vue';

export default {
  name: 'ResourceList',
  components: {
    NavMenu,
  },
  data() {
    return {
      resources: [],
      search: '',
      currentPage: 1, // Đặt trang khởi đầu là 1
      itemsPerPage: 10,
      totalItems: 0,
    };
  },
  computed: {
    filteredResources() {
      return this.resources.filter(r =>
        r.name.toLowerCase().includes(this.search.toLowerCase())
      );
    },
    paginatedResources() {
      const start = 0;
      const end = start + this.itemsPerPage;
      return this.filteredResources.slice(start, end);
    },
    totalPages() {
      return Math.ceil(this.totalItems / this.itemsPerPage);
    },
  },
  mounted() {
    this.fetchResources();
  },
  methods: {
    async fetchResources() {
      try {
        const params = {
          limit: this.itemsPerPage,
          page: this.currentPage,
          search: this.search,
        };
        const res = await axios.get('http://localhost:9999/resources/', { params });
        this.resources = res.data.result || [];
        this.totalItems = res.data.pagingData.total || 0; // Lấy từ pagingData
          // Kiểm tra giá trị totalItems
      } catch (err) {
        console.error('Lỗi khi lấy danh sách resources:', err);
      }
    },
    goToCreate() {
      this.$router.push('/resources/create');
    },
    editResource(id) {
      this.$router.push(`/resources/edit/${id}`);
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        this.fetchResources(); // Gọi lại API để lấy dữ liệu cho trang mới
      }
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.fetchResources(); // Gọi lại API để lấy dữ liệu cho trang mới
      }
    },
  },
  watch: {
    search() {
      this.currentPage = 1; // Reset lại trang về 1 khi tìm kiếm
      this.fetchResources();
    },
  },
};
</script>

<style scoped>
.table td {
  vertical-align: middle;
}
input {
  max-width: 300px;
}
</style>