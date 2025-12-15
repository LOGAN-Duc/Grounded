<template>
  <div class="container mt-5">
    <h2 class="text-center mb-4">Tạo Resource Mới</h2>
    <form @submit.prevent="createResource">
      <!-- Tên Resource -->
      <div class="mb-3">
        <label for="name" class="form-label">Tên</label>
        <input type="text" v-model="form.name" class="form-control" id="name" required />
      </div>

      <!-- Mã Resource -->
      <div class="mb-3">
        <label for="code" class="form-label">Mã</label>
        <input type="text" v-model="form.code" class="form-control" id="code" />
      </div>

      <!-- Upload File -->
      <div class="mb-3">
        <label class="form-label">Upload Image</label>
        <input type="file" class="form-control" @change="handleFileUpload" />
        <div v-if="filePreview" class="mt-2">
          <img :src="filePreview" alt="Preview" class="img-thumbnail" style="max-height: 150px;" />
        </div>
      </div>

      <!-- Loại Resource -->
      <div class="mb-3">
        <label for="resourceTypeId" class="form-label">Loại Resource</label>
        <select v-model="form.resourceTypeId" class="form-select" required>
          <option value="" disabled>
            {{ loadingTypes ? 'Đang tải...' : 'Chọn loại resource' }}
          </option>
          <option v-for="type in resourceTypes" :key="type.id" :value="type.id">
            {{ type.name }}
          </option>
        </select>
      </div>

      <!-- Button Submit -->
      <button type="submit" class="btn btn-primary" :disabled="submitting">
        {{ submitting ? 'Đang tạo...' : 'Tạo Mới' }}
      </button>
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
      file: null,
      filePreview: null,
      resourceTypes: [],
      loadingTypes: false,
      submitting: false,
    };
  },
  created() {
    this.fetchResourceTypes();
  },
  methods: {
    async fetchResourceTypes() {
      this.loadingTypes = true;
      try {
        const res = await axios.get('http://localhost:9999/resources-type/');
        this.resourceTypes = res.data.result || res.data;
      } catch (err) {
        console.error('Lỗi khi lấy dữ liệu:', err);
        alert('Lấy loại resource thất bại.');
      } finally {
        this.loadingTypes = false;
      }
    },
    handleFileUpload(event) {
      const file = event.target.files[0];
      if (file) {
        this.file = file;
        this.filePreview = URL.createObjectURL(file);
      } else {
        this.file = null;
        this.filePreview = null;
      }
    },
    async createResource() {
      if (!this.form.name || !this.form.resourceTypeId) {
        alert('Vui lòng điền đầy đủ thông tin.');
        return;
      }

      this.submitting = true;
      try {
        const formData = new FormData();
        formData.append('name', this.form.name);
        formData.append('code', this.form.code);
        formData.append('resourceTypeId', Number(this.form.resourceTypeId));
        if (this.file) {
          formData.append('file', this.file); // field name 'file' phải trùng backend
        }

        await axios.post('http://localhost:9999/resources/', formData, {
          headers: { 'Content-Type': 'multipart/form-data' },
        });

        alert('Tạo resource thành công!');
        this.$router.push('/resources');
      } catch (err) {
        console.error('Lỗi khi tạo resource:', err);
        alert('Tạo resource thất bại.');
      } finally {
        this.submitting = false;
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
