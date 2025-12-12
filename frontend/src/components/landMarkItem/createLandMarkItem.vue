<template>
  <div class="container mt-4">
    <NavMenu />

    <h2 class="mb-3">Thêm Item & Resource cho Landmark ID: {{ landmarkId }}</h2>

    <div class="row">

      <!-- LEFT: ITEMS -->
      <div class="col-md-6">
        <div class="card shadow-sm mb-4">
          <div class="card-header bg-primary text-white fw-bold">Thêm Item</div>
          <div class="card-body">

            <!-- Search -->
            <div class="mb-3">
              <label class="form-label">Tìm Item</label>
              <input
                type="text"
                class="form-control"
                v-model="itemSearch"
                placeholder="Tìm theo tên item..."
                @input="fetchItems"
              />
            </div>

            <!-- Item list -->
            <div class="scroll-list">
              <ul class="list-group">
                <li
                  v-for="item in limitedItems"
                  :key="item.id"
                  class="list-group-item d-flex align-items-center gap-2 pointer"
                  @click="selectItem(item)"
                >
                  <img v-if="item.urlImage" :src="item.urlImage" class="thumb" />
                  <span>{{ item.name }}</span>
                </li>

                <li v-if="filteredItems.length > maxList">
                  <small class="text-muted d-block p-2 text-center">
                    Chỉ hiển thị {{ maxList }} kết quả — hãy nhập từ khóa để thu hẹp.
                  </small>
                </li>
              </ul>
            </div>

            <!-- Selected Items -->
            <h6 class="mt-3 fw-bold">Items đã chọn</h6>
            <table class="table table-bordered">
              <thead class="table-light">
                <tr>
                  <th>Ảnh</th>
                  <th>Tên Item</th>
                  <th>Hành động</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in selectedItems" :key="item.id">
                  <td><img v-if="item.urlImage" :src="item.urlImage" class="thumb" /></td>
                  <td>{{ item.name }}</td>
                  <td><button class="btn btn-danger btn-sm" @click="removeItem(item.id)">Xóa</button></td>
                </tr>

                <tr v-if="selectedItems.length === 0">
                  <td colspan="3" class="text-center text-muted">Chưa có item nào</td>
                </tr>
              </tbody>
            </table>

          </div>
        </div>
      </div>

      <!-- RIGHT: RESOURCES -->
      <div class="col-md-6">
        <div class="card shadow-sm mb-4">
          <div class="card-header bg-success text-white fw-bold">Thêm Resource</div>
          <div class="card-body">

            <!-- Search -->
            <div class="mb-3">
              <label class="form-label">Tìm Resource</label>
              <input
                type="text"
                class="form-control"
                v-model="resourceSearch"
                placeholder="Tìm theo tên resource..."
                @input="fetchResources"
              />
            </div>

            <!-- Resource list -->
            <div class="scroll-list">
              <ul class="list-group">
                <li
                  v-for="resource in limitedResources"
                  :key="resource.id"
                  class="list-group-item d-flex align-items-center gap-2 pointer"
                  @click="selectResource(resource)"
                >
                  <img v-if="resource.urlImage" :src="resource.urlImage" class="thumb" />
                  <span>{{ resource.name }}</span>
                </li>

                <li v-if="filteredResources.length > maxList">
                  <small class="text-muted d-block p-2 text-center">
                    Chỉ hiển thị {{ maxList }} kết quả — hãy nhập từ khóa để thu hẹp.
                  </small>
                </li>
              </ul>
            </div>

            <!-- Selected Resources -->
            <h6 class="mt-3 fw-bold">Resources đã chọn</h6>
            <table class="table table-bordered">
              <thead class="table-light">
                <tr>
                  <th>Ảnh</th>
                  <th>Tên Resource</th>
                  <th>Hành động</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="resource in selectedResources" :key="resource.id">
                  <td><img v-if="resource.urlImage" :src="resource.urlImage" class="thumb" /></td>
                  <td>{{ resource.name }}</td>
                  <td><button class="btn btn-danger btn-sm" @click="removeResource(resource.id)">Xóa</button></td>
                </tr>

                <tr v-if="selectedResources.length === 0">
                  <td colspan="3" class="text-center text-muted">Chưa có resource nào</td>
                </tr>
              </tbody>
            </table>

          </div>
        </div>
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
      maxList: 50       // Giới hạn hiển thị danh sách dài
    };
  },

  computed: {
    filteredItems() {
      return (this.items || []).filter(item =>
        item.name?.toLowerCase().includes(this.itemSearch.toLowerCase())
      );
    },
    limitedItems() {
      return this.filteredItems.slice(0, this.maxList);
    },

    filteredResources() {
      return (this.resources || []).filter(resource =>
        resource.name?.toLowerCase().includes(this.resourceSearch.toLowerCase())
      );
    },
    limitedResources() {
      return this.filteredResources.slice(0, this.maxList);
    }
  },

  methods: {
    fetchItems() {
      axios.get(`http://localhost:9999/items/`)
        .then(res => {
          const raw = res.data.result || res.data || [];
          this.items = raw.map(item => ({
            ...item,
            urlImage: item.urlImage ? "/items/" + item.urlImage.split('/').pop() : null
          }));
        })
        .catch(console.error);
    },

    fetchResources() {
      axios.get(`http://localhost:9999/resources/`)
        .then(res => {
          const raw = res.data.result || res.data || [];
          this.resources = raw.map(resource => ({
            ...resource,
            urlImage: resource.urlImage ? "/items/" + resource.urlImage.split('/').pop() : null
          }));
        })
        .catch(console.error);
    },

    selectItem(item) {
      if (!this.selectedItems.some(i => i.id === item.id)) {
        this.selectedItems.push(item);
      }
    },

    removeItem(id) {
      this.selectedItems = this.selectedItems.filter(i => i.id !== id);
    },

    selectResource(resource) {
      if (!this.selectedResources.some(r => r.id === resource.id)) {
        this.selectedResources.push(resource);
      }
    },

    removeResource(id) {
      this.selectedResources = this.selectedResources.filter(r => r.id !== id);
    },

    submit() {
  const itemIds = this.selectedItems.map(i => i.id);
  const resourceIds = this.selectedResources.map(r => r.id);

  axios.post(`http://localhost:9999/landmarks/${this.landmarkId}/add-items-resources`, {
    item_ids: itemIds,
    resource_ids: resourceIds
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
.thumb {
  width: 45px;
  height: 45px;
  object-fit: cover;
  border-radius: 6px;
}
.pointer {
  cursor: pointer;
}
.scroll-list {
  max-height: 260px;
  overflow-y: auto;
  border: 1px solid #eee;
  border-radius: 6px;
}
</style>
