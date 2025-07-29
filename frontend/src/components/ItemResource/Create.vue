<template>
  <div class="container mt-4">
    <h2>Tạo Item với ID: {{ item.id }}</h2>
    
    <div class="row">
      <div class="col-md-6">
        <h5>Danh sách dự kiến</h5>
        <table class="table table-bordered">
          <thead>
            <tr>
              <th>Resource</th>
              <th>Số lượng</th>
              <th>Hành động</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="resource in selectedResources" :key="resource.id">
              <td>{{ resource.name }}</td>
              <td>
                <input
                  type="number"
                  v-model="resourceQuantities[resource.id]"
                  min="1"
                  style="width: 70px;"
                  @change="updateQuantity(resource.id)"
                />
              </td>
              <td>
                <button
                  class="btn btn-sm btn-danger"
                  @click="removeResource(resource.id)"
                >
                  Xóa
                </button>
              </td>
            </tr>
            <tr v-if="selectedResources.length === 0">
              <td colspan="3" class="text-center">Chưa có resource nào được thêm</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="col-md-6">
        <div class="mb-3">
          <label for="resourceSearch" class="form-label">Tìm Resource</label>
          <input
            type="text"
            class="form-control"
            id="resourceSearch"
            v-model="resourceSearch"
            placeholder="Tìm theo tên resource..."
            @input="fetchResources"
          />
        </div>

        <h5>Danh sách Resources</h5>
        <ul class="list-group" style="max-height: 300px; overflow-y: auto;">
          <li
            class="list-group-item"
            v-for="resource in filteredResources"
            :key="resource.id"
            @click="selectResource(resource)"
            style="cursor: pointer;"
          >
            {{ resource.name }}
          </li>
        </ul>
      </div>
    </div>

    <button class="btn btn-primary mt-3" @click="submit">Thêm Resource vào Item</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'CreateItem',
  data() {
    return {
      item: {
        id: this.$route.params.id,
      },
      resourceSearch: '',
      resources: [],
      selectedResources: [],
      resourceQuantities: {},
      message: '',
    };
  },
  computed: {
    filteredResources() {
      return this.resources.filter(resource =>
        resource.name.toLowerCase().includes(this.resourceSearch.toLowerCase())
      );
    },
  },
  methods: {
    fetchResources() {
      axios
        .get(`http://localhost:9999/resources/no_item/${this.item.id}`)
        .then(res => {
          this.resources = res.data.result || res.data;
        })
        .catch(err => {
          console.error('Lỗi khi lấy resources:', err);
        });
    },
    selectResource(resource) {
      const existingResource = this.selectedResources.find(r => r.id === resource.id);
      if (existingResource) {
        this.message = 'Resource này đã tồn tại trong danh sách dự kiến. Vui lòng nhập số lượng.';
      } else {
        this.selectedResources.push(resource);
        this.resourceQuantities[resource.id] = 1; // Đặt số lượng mặc định
        this.message = ''; // Reset thông báo
      }
    },
    updateQuantity(resourceId) {
      const quantity = this.resourceQuantities[resourceId];
      if (quantity <= 0) {
        this.resourceQuantities[resourceId] = 1; // Đặt lại số lượng tối thiểu
      }
    },
    removeResource(resourceId) {
      this.selectedResources = this.selectedResources.filter(r => r.id !== resourceId);
      delete this.resourceQuantities[resourceId]; // Xóa số lượng khi xóa resource
    },
    submit() {
  const itemResources = this.selectedResources.map(resource => ({
    resource_id: resource.id,
    quantity: this.resourceQuantities[resource.id],
  }));

  axios
    .post(`http://localhost:9999/item-resources/${this.item.id}`, itemResources)
    .then(() => {
      alert('Thêm resource vào item thành công!');
      this.$router.push('/'); // Quay lại danh sách items
    })
    .catch(err => {
      // Kiểm tra xem phản hồi có chứa thông tin lỗi không
      if (err.response && err.response.data) {
        const errorMessages = err.response.data.errors || [];
        const errorMessage = errorMessages.map((error, index) => 
          `Resource ID ${itemResources[index].resource_id}: ${error.message || error}`
        ).join('\n');
        alert(`Thêm resource thất bại với các lỗi sau:\n${errorMessage}`);
      } else {
        console.error('Lỗi khi thêm resource:', err);
        alert('Thêm resource thất bại.');
      }
    });
}
  },
  mounted() {
    this.fetchResources();
  },
};
</script>

<style scoped>
.table td {
  vertical-align: middle;
}
</style>