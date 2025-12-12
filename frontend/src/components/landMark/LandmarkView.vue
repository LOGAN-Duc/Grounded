<template>
  <div class="container mt-4">
    <NavMenu />

    <div class="d-flex justify-content-between align-items-center mb-3">
      <h2>Items & Resources của Landmark ID: {{ landmarkId }}</h2>
      <input
        type="text"
        class="form-control w-50"
        placeholder="Tìm item/resource..."
        v-model="search"
      />
    </div>

    <h4>Items</h4>
    <table class="table table-bordered align-middle" v-if="items.length">
      <thead> 
        <tr>
          <th>Ảnh</th>
          <th>Tên Item</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in filteredItems" :key="item.id">
          <td>
            <img v-if="item.urlImage" :src="item.urlImage" class="thumb" />
            <span v-else>-</span>
          </td>
          <td>{{ item.name || '-' }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else class="text-center text-muted">Chưa có Item nào</div>

    <h4 class="mt-4">Resources</h4>
    <table class="table table-bordered align-middle" v-if="resources.length">
      <thead>
        <tr>
          <th>Ảnh</th>
          <th>Tên Resource</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="res in filteredResources" :key="res.id">
          <td>
            <img v-if="res.urlImage" :src="res.urlImage" class="thumb" />
            <span v-else>-</span>
          </td>
          <td>{{ res.name || '-' }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else class="text-center text-muted">Chưa có Resource nào</div>
  </div>
</template>

<script>
import axios from 'axios';
import NavMenu from '../layout/NavMenu.vue';

export default {
  name: 'LandmarkView',
  components: { NavMenu },
  data() {
    return {
      landmarkId: this.$route.params.id,
      items: [],
      resources: [],
      search: ''
    };
  },
  computed: {
    filteredItems() {
      const search = this.search.toLowerCase();
      return this.items.filter(i => i.name && i.name.toLowerCase().includes(search));
    },
    filteredResources() {
      const search = this.search.toLowerCase();
      return this.resources.filter(r => r.name && r.name.toLowerCase().includes(search));
    }
  },
  methods: {
    async fetchItemsResources() {
      try {
        const res = await axios.get('http://localhost:9999/landmark-items/', {
          params: { landmark_id: this.landmarkId }
        });
        const data = res.data.data || [];

        this.items = data
          .filter(d => d.item_id)
          .map(d => ({
            id: d.item_id,
            name: d.item?.name || d.item_name || 'Unknown',
            urlImage: d.item?.urlImage
              ? '/items/' + d.item.urlImage.split('/').pop()
              : d.item_url_image
                ? '/items/' + d.item_url_image.split('/').pop()
                : null
          }));

        this.resources = data
          .filter(d => d.resource_id)
          .map(d => ({
            id: d.resource_id,
            name: d.resource?.name || d.resource_name || 'Unknown',
            urlImage: d.resource?.urlImage
              ? '/items/' + d.resource.urlImage.split('/').pop()
              : d.resource_url_image
                ? '/items/' + d.resource_url_image.split('/').pop()
                : null
          }));
      } catch (err) {
        console.error(err);
        alert('Lấy dữ liệu thất bại');
      }
    }
  },
  mounted() {
    this.fetchItemsResources();
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
</style>
