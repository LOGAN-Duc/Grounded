<template>
  <div class="container mt-4">
    <NavMenu />

    <h2>Thêm Item & Resource cho Landmark ID: {{ landmarkId }}</h2>

    <div class="row mt-3">
      <!-- Items -->
      <div class="col-md-6">
        <div class="mb-3">
          <label for="itemSearch" class="form-label">Tìm Item</label>
          <input
            type="text"
            class="form-control"
            id="itemSearch"
            v-model="itemSearch"
            placeholder="Tìm theo tên item..."
            @input="fetchItems"
          />
        </div>

        <h5>Danh sách Items</h5>
        <ul class="list-group" style="max-height: 250px; overflow-y: auto;">
          <li
            class="list-group-item d-flex align-items-center"
            v-for="item in filteredItems"
            :key="item.id"
            @click="selectItem(item)"
            style="cursor: pointer;"
          >
            <span>{{ item.name }}</span>
          </li>
        </ul>

        <h6 class="mt-3">Items đã chọn</h6>
        <table class="table table-bordered">
          <thead>
            <tr>
              <th>Tên Item</th>
              <th>Số lượng</th>
              <th>Hành động</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in selectedItems" :key="item.id">
              <td>{{ item.name }}</td>
              <td>
                <input
                  type="number"
                  v-model.number="itemQuantities[item.id]"
                  min="1"
                  style="width: 70px;"
                />
              </td>
              <td>
                <button class="btn btn-sm btn-danger" @click="removeItem(item.id)">Xóa</button>
              </td>
            </tr>
            <tr v-if="selectedItems.length === 0">
              <td colspan="3" class="text-center">Chưa có item nào được thêm</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Resources -->
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
        <ul class="list-group" style="max-height: 250px; overflow-y: auto;">
          <li
            class="list-group-item d-flex align-items-center"
            v-for="resource in filteredResources"
            :key="resource.id"
            @click="selectResource(resource)"
            style="cursor: pointer;"
          >
            <span>{{ resource.name }}</span>
          </li>
        </ul>

        <h6 class="mt-3">Resources đã chọn</h6>
        <table class="table table-bordered">
          <thead>
            <tr>
              <th>Tên Resource</th>
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
                  v-model.number="resourceQuantities[resource.id]"
                  min="1"
                  style="width: 70px;"
                />
              </td>
              <td>
                <button class="btn btn-sm btn-danger" @click="removeResource(resource.id)">Xóa</button>
              </td>
            </tr>
            <tr v-if="selectedResources.length === 0">
              <td colspan="3" class="text-center">Chưa có resource nào được thêm</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <button class="btn btn-primary mt-3" @click="submit">Thêm Item & Resource</button>
  </div>
</template>

<script>
import axios from 'axios';
import NavMenu from '../layout/NavMenu.vue';

export default {
  name: 'CreateLandmarkItem',
  components: { NavMenu },
  data() {
    return {
      landmarkId: this.$route.params.id,
      itemSearch: '',
      resourceSearch: '',
      items: [],
      resources: [],
      selectedItems: [],
      selectedResources: [],
      itemQuantities: {},
      resourceQuantities: {},
    };
  },
  computed: {
    filteredItems() {
      return this.items.filter(item =>
        item.name.toLowerCase().includes(this.itemSearch.toLowerCase())
      );
    },
    filteredResources() {
      return this.resources.filter(resource =>
        resource.name.toLowerCase().includes(this.resourceSearch.toLowerCase())
      );
    },
  },
  methods: {
    fetchItems() {
      axios.get(`http://localhost:9999/items/no_landmark/${this.landmarkId}`)
        .then(res => { this.items = res.data.result || []; })
        .catch(err => console.error(err));
    },
    fetchResources() {
      axios.get(`http://localhost:9999/resources/no_landmark/${this.landmarkId}`)
        .then(res => { this.resources = res.data.result || []; })
        .catch(err => console.error(err));
    },
    selectItem(item) {
      if (!this.selectedItems.find(i => i.id === item.id)) {
        this.selectedItems.push(item);
        this.itemQuantities[item.id] = 1;
      }
    },
    removeItem(itemId) {
      this.selectedItems = this.selectedItems.filter(i => i.id !== itemId);
      delete this.itemQuantities[itemId];
    },
    selectResource(resource) {
      if (!this.selectedResources.find(r => r.id === resource.id)) {
        this.selectedResources.push(resource);
        this.resourceQuantities[resource.id] = 1;
      }
    },
    removeResource(resourceId) {
      this.selectedResources = this.selectedResources.filter(r => r.id !== resourceId);
      delete this.resourceQuantities[resourceId];
    },
    submit() {
      const itemsPayload = this.selectedItems.map(i => ({
        item_id: i.id,
        quantity: this.itemQuantities[i.id] || 1
      }));
      const resourcesPayload = this.selectedResources.map(r => ({
        resource_id: r.id,
        quantity: this.resourceQuantities[r.id] || 1
      }));

      axios.post(`http://localhost:9999/landmarks/${this.landmarkId}/add-items-resources`, {
        items: itemsPayload,
        resources: resourcesPayload
      })
      .then(() => {
        alert('Thêm thành công!');
        this.$router.push('/landmarks');
      })
      .catch(err => {
        console.error(err);
        alert('Thêm thất bại!');
      });
    }
  },
  mounted() {
    this.fetchItems();
    this.fetchResources();
  }
};
</script>

<style scoped>
.table td { vertical-align: middle; }
.list-group-item:hover { background-color: #f1f1f1; cursor: pointer; }
</style>
