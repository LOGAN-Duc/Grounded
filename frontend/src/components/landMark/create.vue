<template>
  <div class="container mt-4">
    <NavMenu />

    <h2 class="mb-4">Tạo Landmark</h2>

    <form @submit.prevent="createLandmark">
      <div class="mb-3">
        <label for="name" class="form-label">Tên Landmark</label>
        <input type="text" id="name" v-model="form.name" class="form-control" required />
      </div>

      <div class="mb-3">
        <label for="description" class="form-label">Mô tả</label>
        <textarea id="description" v-model="form.description" class="form-control" rows="3"></textarea>
      </div>

      <div class="mb-3">
        <label for="image" class="form-label">Ảnh</label>
        <input type="file" id="image" @change="onFileChange" class="form-control" accept="image/*" />
        <div v-if="previewImage" class="mt-2">
          <img :src="previewImage" alt="preview" style="width:120px;height:120px;object-fit:cover;" />
        </div>
      </div>

      <button type="submit" class="btn btn-primary" :disabled="loading">
        {{ loading ? 'Đang tạo...' : 'Tạo Landmark' }}
      </button>
      <button type="button" class="btn btn-secondary ms-2" @click="goBack">Hủy</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import NavMenu from '../layout/NavMenu.vue';

export default {
  name: 'LandmarkCreate',
  components: { NavMenu },
  data() {
    return {
      form: {
        name: '',
        description: '',
        image: null, // file image
      },
      previewImage: null,
      loading: false,
    };
  },
  methods: {
    // Khi chọn file → lưu vào form.image và tạo preview
    onFileChange(event) {
      const file = event.target.files[0];
      if (file) {
        this.form.image = file;
        this.previewImage = URL.createObjectURL(file);
      }
    },

    async createLandmark() {
      if (!this.form.name) return alert('Tên Landmark là bắt buộc');
      this.loading = true;
      try {
        const formData = new FormData();
        formData.append('name', this.form.name);
        formData.append('description', this.form.description || '');

        // đổi field 'image' thành 'file' để backend hiện tại nhận đúng
        if (this.form.image) formData.append('file', this.form.image);

        // Gửi POST
        const res = await axios.post('http://localhost:9999/landmarks/', formData);

        console.log('Create response:', res.data);
        alert('Tạo Landmark thành công');
        this.$router.push('/landmarks');
      } catch (err) {
        console.error('Create error:', err.response || err);
        alert('Tạo Landmark thất bại: ' + (err.response?.data?.message || err.message));
      } finally {
        this.loading = false;
      }
    },

    goBack() {
      this.$router.push('/landmarks');
    }
  }
};
</script>
