<template>
  <div class="container mt-4">
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

    <table class="table table-bordered align-middle">
      <thead class="thead-light">
        <tr>
          <th style="white-space: nowrap; width: 200px;">Name</th>
          <th style="width: 200px;">Code</th>
          <th style="width: 100px;">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in filteredItems" :key="item.id">
          <td style="white-space: nowrap; width: 200px;">{{ item.name }}</td>
          <td class="text-truncate position-relative" style="width: 200px;">
            <div class="d-flex justify-content-between align-items-center">
              <span
                class="text-truncate"
                style="width: 200px;"
                :title="item.code"
              >
                {{ item.code }}
              </span>
              <button
                class="btn btn-sm btn-outline-secondary"
                @click="copyToClipboard(item.code)"
              >
                Copy
              </button>
            </div>
          </td>
          <td>
            <button
              class="btn btn-sm btn-warning"
              @click="updateItem(item.id)"
            >
              Update
            </button>
          </td>
        </tr>
        <tr v-if="filteredItems.length === 0">
          <td colspan="3" class="text-center">Không tìm thấy item nào</td>
        </tr>
      </tbody>
    </table>

    <div class="mt-3">
      <button class="btn btn-primary" @click="goToCreate">+ Create Item</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ItemList',
  data() {
    return {
      items: [],
      search: '',
    };
  },
  computed: {
    filteredItems() {
      return this.items.filter((item) =>
        item.name.toLowerCase().includes(this.search.toLowerCase())
      );
    },
  },
  methods: {
    copyToClipboard(text) {
      navigator.clipboard.writeText(text).then(() => {
        alert('Đã sao chép!');
      });
    },
    updateItem(id) {
      const item = this.items.find((i) => i.id === id);
      if (!item) {
        alert('Không tìm thấy item để cập nhật.');
        return;
      }

      const updateData = {
        name: item.name,
        code: '', // đặt lại code là chuỗi rỗng như yêu cầu
      };

      axios
        .put(`http://localhost:9999/items/${id}`, updateData)
        .then(() => {
          alert('Cập nhật thành công!');
          this.fetchItems(); // reload danh sách nếu cần
        })
        .catch((err) => {
          console.error('Lỗi khi cập nhật:', err.response?.data || err.message);
          alert('Cập nhật thất bại.');
        });
    },
    goToCreate() {
      this.$router.push('/create');
    },
    fetchItems() {
      axios
        .get('http://localhost:9999/items/')
        .then((res) => {
          this.items = res.data.result || res.data;
        })
        .catch((err) => {
          console.error('Lỗi khi lấy dữ liệu:', err);
        });
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
</style>
