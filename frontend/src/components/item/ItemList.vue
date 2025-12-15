<template>
  <div class="container mt-4">
    <NavMenu />

    <!-- HEADER -->
    <div class="d-flex justify-content-between align-items-center mb-3">
      <h2 class="mb-0">Danh sách Items</h2>

      <div class="d-flex gap-2 w-50 align-items-center">
        <input
          type="text"
          class="form-control"
          placeholder="Tìm theo tên..."
          v-model="search"
        />

        <select class="form-select" v-model="selectedType">
          <option value="">-- Tất cả loại --</option>
          <option v-for="type in itemTypes" :key="type.id" :value="type.name">
            {{ type.name }}
          </option>
        </select>
      </div>
    </div>

    <div class="mb-3">
      <button class="btn btn-primary" @click="goToCreate">+ Tạo Item</button>
    </div>

    <div v-if="loading" class="text-center py-4">Đang tải dữ liệu...</div>

    <!-- TABLE -->
    <table class="table table-bordered table-hover align-middle" v-else>
      <thead class="table-light">
        <tr>
          <th style="width: 180px;">Name</th>
          <th style="width: 120px;">Image</th>
          <th style="width: 120px;">Loại Item</th>
          <th style="width: 160px;">Code</th>
          <th style="width: 220px;">Actions</th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="item in items"
          :key="item.id"
          @click="selectItem(item)"
          :class="{ 'table-active': selectedItem?.id === item.id }"
          style="cursor: pointer;"
        >
          <td>{{ item.name }}</td>

          <td>
            <img
              v-if="item.urlImage"
              :src="item.urlImage"
              class="item-image"
            />
            <span v-else class="text-muted">Không có ảnh</span>
          </td>

          <td>{{ item.item_Type?.name || 'N/A' }}</td>

          <!-- CODE -->
          <td>
            <div class="code-cell">
              <span class="code-text" :title="item.code">
                {{ item.code }}
              </span>
              <button
                class="btn btn-sm btn-outline-secondary"
                @click.stop="copyToClipboard(item.code)"
              >
                Copy
              </button>
            </div>
          </td>

          <!-- ACTIONS -->
          <td>
            <div class="d-flex gap-1 flex-wrap">
              <button
                class="btn btn-sm btn-warning"
                @click.stop="updateItem(item.id)"
              >
                Cập nhật
              </button>
              <button
                class="btn btn-sm btn-info"
                @click.stop="edit(item.id)"
              >
                Sửa
              </button>
              <button
                class="btn btn-sm btn-success"
                @click.stop="addResource(item.id)"
              >
                Resource
              </button>
            </div>
          </td>
        </tr>

        <tr v-if="items.length === 0">
          <td colspan="5" class="text-center text-muted">
            Không có dữ liệu
          </td>
        </tr>
      </tbody>
    </table>

    <!-- PAGINATION -->
    <div class="d-flex justify-content-between align-items-center mt-3">
      <button
        class="btn btn-secondary"
        @click="prevPage"
        :disabled="currentPage === 1"
      >
        Trước
      </button>

      <span>Trang {{ currentPage }}</span>

      <button class="btn btn-secondary" @click="nextPage">
        Sau
      </button>
    </div>

    <!-- GENERATE CODE -->
    <div v-if="selectedItem" class="mt-4">
      <h5>Đã chọn: {{ selectedItem.name }}</h5>

      <div class="d-flex align-items-center gap-2 mb-2">
        <label class="mb-0">Số lượng:</label>
        <input
          type="number"
          min="1"
          v-model.number="selectedNumber"
          class="form-control w-auto"
        />
      </div>

      <textarea
        class="form-control"
        rows="5"
        readonly
        :value="generatedCodes.join('\n')"
      ></textarea>

      <button class="btn btn-outline-secondary mt-2" @click="copyCodes">
        Copy tất cả
      </button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import NavMenu from '../layout/NavMenu.vue'

export default {
  name: 'ItemList',
  components: { NavMenu },

  data() {
    return {
      items: [],
      itemTypes: [],
      selectedType: '',
      search: '',
      loading: false,
      currentPage: 1,
      itemsPerPage: 10,
      totalItems: 0,
      selectedItem: null,
      selectedNumber: 1,
    }
  },

  computed: {
    generatedCodes() {
      if (!this.selectedItem) return []
      return Array.from(
        { length: this.selectedNumber },
        () => this.selectedItem.code
      )
    },
  },

  methods: {
    async fetchItemTypes() {
      const res = await axios.get('/items-type')
      this.itemTypes = res.data.result || []
    },

    async fetchItems() {
      this.loading = true
      try {
        const res = await axios.get('/items', {
          params: {
            page: this.currentPage,
            limit: this.itemsPerPage,
            search: this.search,
            type_Item: this.selectedType,
          },
        })

        this.items = (res.data.result || []).map(i => {
          if (i.urlImage) {
            i.urlImage = '/items/' + i.urlImage.split('/').pop()
          }
          return i
        })

        this.totalItems = res.data.pagingData?.total || 0
      } finally {
        this.loading = false
      }
    },

    copyToClipboard(text) {
      navigator.clipboard.writeText(text)
      alert('Đã sao chép')
    },

    copyCodes() {
      navigator.clipboard.writeText(this.generatedCodes.join('\n'))
      alert('Đã copy toàn bộ')
    },

    edit(id) {
      this.$router.push(`/items/edit/${id}`)
    },

    addResource(id) {
      this.$router.push(`/itemresource/create/${id}`)
    },

    updateItem(id) {
      const item = this.items.find(i => i.id === id)
      axios.put(`/items/${id}`, {
        name: item.name,
        code: item.code,
      })
    },

    goToCreate() {
      this.$router.push('/items/create')
    },

    nextPage() {
      this.currentPage++
      this.fetchItems()
    },

    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--
        this.fetchItems()
      }
    },

    selectItem(item) {
      this.selectedItem = item
      this.selectedNumber = 1
    },
  },

  watch: {
    search() {
      this.currentPage = 1
      this.fetchItems()
    },
    selectedType() {
      this.currentPage = 1
      this.fetchItems()
    },
  },

  mounted() {
    this.fetchItemTypes()
    this.fetchItems()
  },
}
</script>

<style scoped>
.item-image {
  width: 100px;
  height: 100px;
  object-fit: cover;
  border-radius: 6px;
}

.code-cell {
  display: flex;
  align-items: center;
  gap: 4px;
}

.code-text {
  max-width: 110px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.table td {
  vertical-align: middle;
}

.table-active {
  background-color: #e6f7ff !important;
}
</style>
