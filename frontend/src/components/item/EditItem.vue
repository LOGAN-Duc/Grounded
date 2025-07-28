<template>
  <div class="container mt-4">
    <h2>Edit Item</h2>
    <form @submit.prevent="updateItem">
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input
          type="text"
          class="form-control"
          v-model="item.name"
          required
        />
      </div>
      <div class="mb-3">
        <label for="code" class="form-label">Code</label>
        <input
          type="text"
          class="form-control"
          v-model="item.code"
          required
        />
      </div>

      <h5>Danh sách Resources</h5>
      <table class="table table-bordered align-middle">
        <thead class="thead-light">
          <tr>
            <th style="white-space: nowrap; width: 200px;">Tên</th>
            <th style="width: 100px;">Mã</th>
            <th style="width: 50px;">Hành động</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="resource in resources" :key="resource.id">
            <td style="white-space: nowrap;">{{ resource.name }}</td>
            <td :title="resource.code">{{ resource.code.substring(0, 10) }}...</td>
          </tr>
        </tbody>
      </table>

      <button type="submit" class="btn btn-primary">Save</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ItemEdit',
  props: ['id'],
  data() {
    return {
      item: {
        name: '',
        code: '',
      },
      resources: [],
    };
  },
  created() {
    this.fetchItem();
    this.fetchResources();
  },
  methods: {
    fetchItem() {
      axios
        .get(`http://localhost:9999/items/${this.id}`)
        .then((res) => {
          this.item = res.data; // Giả sử API trả về item đầy đủ
        })
        .catch((err) => {
          console.error('Lỗi khi lấy item:', err);
        });
    },
    fetchResources() {
      axios
        .get('http://localhost:9999/resources/no_item')
        .then(res => {
          this.resources = res.data.result || res.data;
        })
        .catch(err => {
          console.error('Lỗi khi lấy resources:', err);
        });
    },
    updateItem() {
      axios
        .put(`http://localhost:9999/items/${this.id}`, this.item)
        .then(() => {
          alert('Cập nhật thành công!');
          this.$router.push('/'); // Quay lại danh sách items
        })
        .catch((err) => {
          console.error('Lỗi khi cập nhật:', err);
          alert('Cập nhật thất bại.');
        });
    },
    // removeResource(resourceId) {
    //   axios
    //     .delete(`http://localhost:9999/item-resources/${this.id}/${resourceId}`)
    //     .then(() => {
    //       this.resources = this.resources.filter((r) => r.id !== resourceId);
    //       alert('Resource đã được xóa!');
    //     })
    //     .catch((err) => {
    //       console.error('Lỗi khi xóa resource:', err);
    //       alert('Xóa resource thất bại.');
    //     });
    // },
  },
};
</script>

<style scoped>
/* Thêm style nếu cần */
</style>