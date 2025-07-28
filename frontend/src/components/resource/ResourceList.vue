<template>
  <div class="container mt-4">
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
        <tr v-for="resource in filteredResources" :key="resource.id">
          <td style="white-space: nowrap;">{{ resource.name }}</td>
          <td :title="resource.code">{{ resource.code.substring(0, 10) }}...</td>
          <td>{{ resource.resourceType.name || 'Không xác định' }}</td>
          <td>
            <button class="btn btn-sm btn-warning" @click="editResource(resource.id)">Edit</button>
          </td>
        </tr>
        <tr v-if="filteredResources.length === 0">
          <td colspan="4" class="text-center">Không tìm thấy resource nào</td>
        </tr>
      </tbody>
    </table>

    <div class="mt-3">
      <button class="btn btn-primary" @click="goToCreate">+ Tạo Resource</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ResourceList',
  data() {
    return {
      resources: [],
      search: '',
    };
  },
  computed: {
    filteredResources() {
      const keyword = this.search.toLowerCase();
      return this.resources.filter(r =>
        r.name.toLowerCase().includes(keyword)
      );
    },
  },
  mounted() {
    this.fetchResources();
  },
  methods: {
    async fetchResources() {
      try {
        const res = await axios.get('http://localhost:9999/resources/');
        this.resources = res.data.result || [];
      } catch (err) {
        console.error('Lỗi khi lấy danh sách resources:', err);
      }
    },
    goToCreate() {
      this.$router.push('/resources/create'); // Điều hướng đến trang tạo mới resource
    },
    editResource(id) {
      this.$router.push(`/resources/edit/${id}`); // Điều hướng đến trang chỉnh sửa resource
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