<template>
  <div class="container mt-4">
    <h2>Edit Item {{ item.id }}</h2>
    <form @submit.prevent="updateItem">
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input
          type="text"
          class="form-control"
          v-model="item.name"
        />
      </div>
      <div class="mb-3">
        <label for="name" class="form-label">Code</label>
        <input
          type="text"
          class="form-control"
          v-model="item.code"
        />
      </div>
      <div class="mb-3">
        <label for="itemTypeId" class="form-label">Type of Item</label>
        <select v-model="item.itemTypeId" class="form-select">
          <option value="" disabled selected>Chọn loại item</option>
          <option v-for="type in itemTypes" :key="type.id" :value="type.id">
            {{ type.name }}
          </option>
        </select>
      </div>
      <h5>Danh sách Item Resources</h5>
      <table class="table table-bordered align-middle">
        <thead class="thead-light">
          <tr>
            <th style="white-space: nowrap; width: 100px;">Tên</th>
            <th style="width: 100px;">Số lượng</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="itemresource in itemresources" :key="itemresource.id">
            <td style="white-space: nowrap;">{{ itemresource.resource.name }}</td>
            <td>
              <input
                type="number"
                v-model="itemresource.quantity"
                class="form-control"
                min="0"
              required
              />
            </td>
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
        itemTypeId: null,
        id: this.$route.params.id,
      },
       itemTypes: [],
      itemresources: [],
    };
  },
  created() {
    this.fetchItem();
    this.fetchItemResources();
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
    fetchItemResources() {
      axios
        .get(`http://localhost:9999/item-resources/${this.id}`)
        .then(res => {
          this.itemresources = res.data.result || res.data;
        })
        .catch(err => {
          console.error('Lỗi khi lấy item resources:', err);
        });
    },
   updateItem() {
  // Cập nhật item
  axios
    .put(`http://localhost:9999/items/${this.id}`, this.item)
    .then(() => {
      // Cập nhật item resources
      const itemResourceUpdates = this.itemresources.map(itemresource => ({
        resource_id: itemresource.resource.id, // Sử dụng resource.id
        quantity: itemresource.quantity,
        itemTypeId: Number(this.item.itemTypeId), 
      }));

      // Gửi yêu cầu cập nhật item resources
      return axios.put(`http://localhost:9999/item-resources/${this.id}`, itemResourceUpdates);
    })
    .then(() => {
      // Gọi đến yêu cầu cập nhật trạng thái
      return axios.put(`http://localhost:9999/item-resources/status`);
    })
    .then(() => {
      alert('Cập nhật thành công!');
      this.$router.push('/'); // Quay lại danh sách items
    })
    .catch((err) => {
      console.error('Lỗi khi cập nhật:', err);
      alert('Cập nhật thất bại.');
    });
}
  },
};
</script>