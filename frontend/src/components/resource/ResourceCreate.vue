<template>
  <div class="container mt-5">
    <h2 class="text-center mb-4">Tạo Resource Mới</h2>
    <form @submit.prevent="createResource">
      <div class="mb-3">
        <label for="name" class="form-label">Tên</label>
        <input type="text" v-model="form.name" class="form-control" id="name" required />
      </div>
      <div class="mb-3">
        <label for="code" class="form-label">Mã</label>
        <input type="text" v-model="form.code" class="form-control" id="code" required />
      </div>
      <div class="mb-3">
        <label for="resourceTypeId" class="form-label">Loại Resource</label>
        <select v-model="form.resourceTypeId" class="form-select" required>
          <option value="" disabled>Chọn loại resource</option>
          <option v-for="type in resourceTypes" :key="type.id" :value="type.id">
            {{ type.name }}
          </option>
        </select>
      </div>
      <button type="submit" class="btn btn-primary">Tạo Mới</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'CreateResource',
  data() {
    return {
      form: {
        name: '',
        code: '',
        resourceTypeId: '',
      },
      resourceTypes: [],
    };
  },
  created() {
    this.fetchResourceTypes();
  },
  methods: {
     fetchResourceTypes() {
      axios
        .get('http://localhost:9999/resources-type/')
        .then((res) => {
          this.resourceTypes = res.data.result || res.data;
        })
        .catch((err) => {
          console.error('Lỗi khi lấy dữ liệu:', err);
        });
    },
    async createResource() {
      try {
        const payload = {
          ...this.form,
          resourceTypeId: Number(this.form.resourceTypeId), // Ép kiểu về số
        };
        await axios.post('http://localhost:9999/resources/', payload); // Đổi đường dẫn nếu cần
        alert('Tạo resource thành công!');
        this.$router.push('/resources');
      } catch (err) {
        console.error('Lỗi khi tạo resource:', err);
        alert('Tạo resource thất bại.');
      }
    },
  },
};
</script>

<style scoped>
.container {
  max-width: 600px;
}
</style>