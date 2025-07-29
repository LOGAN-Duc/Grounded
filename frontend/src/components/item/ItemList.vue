<template>
  <div class="container mt-4">
    <NavMenu />
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h2 class="mb-0">Danh sách Items</h2>
      <div class="d-flex gap-2 w-50">
        <input
          type="text"
          class="form-control"
          placeholder="Tìm theo tên..."
          v-model="search"
        />
      </div>
    </div>

    <div v-if="loading" class="text-center">Đang tải dữ liệu...</div>

    <table class="table table-bordered align-middle" v-else>
      <thead class="thead-light">
        <tr>
          <th style="white-space: nowrap; width: 200px;">Name</th>
          <th style="width: 100px;">Loại Item</th>
          <th style="width: 200px;">Code</th>
          <th style="width: 200px;">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in paginatedItems" :key="item.id">
          <td style="white-space: nowrap; width: 200px;">{{ item.name }}</td>
          <td>{{ item.item_Type?.name || 'Không xác định' }}</td>
          <td class="text-truncate position-relative" style="width: 200px;">
            <div class="d-flex justify-content-between align-items-center">
              <span class="text-truncate" style="width: 200px;" :title="item.code">
                {{ item.code }}
              </span>
              <button class="btn btn-sm btn-outline-secondary" @click="copyToClipboard(item.code)">Copy</button>
            </div>
          </td>
          <td>
            <button class="btn btn-sm btn-warning btn-custom" @click="updateItem(item.id)">Cập nhật</button>
            <button class="btn btn-sm btn-info btn-custom" @click="edit(item.id)">Chỉnh sửa</button>
            <button class="btn btn-sm btn-success btn-custom" @click="addResource(item.id)">Thêm Resource</button>
          </td>
        </tr>
        <tr v-if="paginatedItems.length === 0">
          <td colspan="4" class="text-center">Không tìm thấy item nào</td>
        </tr>
      </tbody>
    </table>

    <div class="d-flex justify-content-between align-items-center mt-3">
      <button class="btn btn-secondary" @click="prevPage" :disabled="currentPage === 1">Trước</button>
      <span>Trang {{ currentPage }} / {{ totalPages }}</span>
      <button class="btn btn-secondary" @click="nextPage" :disabled="currentPage === totalPages">Sau</button>
    </div>

    <div class="mt-3">
      <button class="btn btn-primary" @click="goToCreate">+ Tạo Item</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import NavMenu from '../layout/NavMenu.vue';

export default {
  name: 'ItemList',
  components: {
    NavMenu,
  },
  data() {
    return {
      items: [],
      search: '',
      loading: false,
      currentPage: 1,
      itemsPerPage: 10,
      totalItems: 0,
    };
  },
  computed: {
    filteredItems() {
      return this.items.filter(item =>
        item.name.toLowerCase().includes(this.search.toLowerCase())
      );
    },
    paginatedItems() {
      const start = 0;
      const end = start + this.itemsPerPage;
      return this.filteredItems.slice(start, end);
    },
    totalPages() {
      return Math.ceil(this.totalItems / this.itemsPerPage);
    },
  },
  methods: {
    async fetchItems() {
      this.loading = true;
      try {
        const params = {
          limit: this.itemsPerPage,
          page: this.currentPage,
          search: this.search,
        };
        const res = await axios.get('http://localhost:9999/items/', { params });
        this.items = res.data.result || [];
        this.totalItems = res.data.pagingData.total || 0;// Giả sử API trả về tổng số items
      } catch (err) {
        console.error('Lỗi khi lấy dữ liệu:', err);
      } finally {
        this.loading = false;
      }
    },
    copyToClipboard(text) {
      navigator.clipboard.writeText(text).then(() => {
        alert('Đã sao chép!');
      });
    },
    edit(id) {
      this.$router.push(`/items/edit/${id}`);
    },
    addResource(itemId) {
      this.$router.push(`/itemresource/create/${itemId}`);
    },
    updateItem(id) {
      const item = this.items.find(i => i.id === id);
      if (!item) {
        alert('Không tìm thấy item để cập nhật.');
        return;
      }

      const updateData = {
        name: item.name,
        code: item.code, // Đặt lại code là chuỗi rỗng như yêu cầu
      };

      axios
        .put(`http://localhost:9999/items/${id}`, updateData)
        .then(() => {
          alert('Cập nhật thành công!');
          this.fetchItems(); // Reload danh sách nếu cần
        })
        .catch(err => {
          console.error('Lỗi khi cập nhật:', err.response?.data || err.message);
          alert('Cập nhật thất bại.');
        });
    },
    goToCreate() {
      this.$router.push('/items/create');
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        this.fetchItems(); // Tải lại dữ liệu khi chuyển trang
      }
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.fetchItems(); // Tải lại dữ liệu khi chuyển trang
      }
    },
  },
  watch: {
    search() {
      this.currentPage = 1; // Reset lại trang về 1 khi tìm kiếm
      this.fetchItems();
    },
  },
  mounted() {
    this.fetchItems();
  },
};
</script>

<style scoped>
.table td {
  vertical-align: middle;
}

.btn-custom {
  transition: background-color 0.3s, color 0.3s;
}

.btn-custom:hover {
  filter: brightness(90%);
}
</style>