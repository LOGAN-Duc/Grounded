<template>
  <div class="container mt-4">
    <h2>Chỉnh sửa Resource</h2>
    <form @submit.prevent="updateResource">
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input
          type="text"
          class="form-control"
          v-model="resource.name"
          required
        />
      </div>
      <div class="mb-3">
        <label for="code" class="form-label">Code</label>
        <input
          type="text"
          class="form-control"
          v-model="resource.code"
          required
        />
      </div>
      <button type="submit" class="btn btn-primary">Cập nhật</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ResourceEdit',
  props: ['id'],
  data() {
    return {
      resource: {
        name: '',
        code: '',
      },
    };
  },
  created() {
    this.fetchResource();
  },
  methods: {
    fetchResource() {
      axios
        .get(`http://localhost:9999/resources/${this.id}`)
        .then((res) => {
          this.resource = res.data; // Giả sử API trả về resource đầy đủ
        })
        .catch((err) => {
          console.error('Lỗi khi lấy resource:', err);
        });
    },
    updateResource() {
      axios
        .put(`http://localhost:9999/resources/${this.id}`, this.resource)
        .then(() => {
          alert('Cập nhật thành công!');
          this.$router.push('/resources'); // Quay lại danh sách resources
        })
        .catch((err) => {
          console.error('Lỗi khi cập nhật:', err);
          alert('Cập nhật thất bại.');
        });
    },
  },
};
</script>

<style scoped>
/* Thêm style nếu cần */
</style>