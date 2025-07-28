<template>
  <div class="container mt-5">
    <h2 class="text-center mb-4">Create Item</h2>
    <form @submit.prevent="createItem">
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input type="text" v-model="form.name" class="form-control" id="name" required />
      </div>
      <div class="mb-3">
        <label for="code" class="form-label">Code</label>
        <input type="text" v-model="form.code" class="form-control" id="code"/>
      </div>
      <div class="mb-3">
        <label for="itemTypeId" class="form-label">Type of Item</label>
        <select v-model="form.itemTypeId" class="form-select" required>
          <option value="" disabled selected>Chọn loại item</option>
          <option v-for="type in itemTypes" :key="type.id" :value="type.id">
            {{ type.name }}
          </option>
        </select>
      </div>
      <button type="submit" class="btn btn-primary">Tạo mới</button>
    </form>

  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'CreateItem',
  
  created() {
    this.fetchItemTypes();
  },
  methods: {
    fetchItemTypes() {
      axios
        .get('http://localhost:9999/items-type/')
        .then((res) => {
          this.itemTypes = res.data.result || res.data;
        })
        .catch((err) => {
          console.error('Lỗi khi lấy dữ liệu:', err);
        });
    },
    async createItem() {
    try {
      const payload = {
        ...this.form,
        itemTypeId: Number(this.form.itemTypeId), // ✅ ép kiểu về số
      };
      await axios.post('http://localhost:9999/items/', payload);
      alert('Tạo item thành công!');
      this.$router.push('/');
    } catch (err) {
      console.error('Lỗi khi tạo item:', err);
      alert('Tạo item thất bại.');
    }
    }

  },
  data() {
    return {
      form: {
        name: '',
        code: '',
        itemTypeId: '',
      },
      itemTypes: [],
    };
  },
};
</script>

<style scoped>
.container {
  max-width: 600px;
}
</style>
