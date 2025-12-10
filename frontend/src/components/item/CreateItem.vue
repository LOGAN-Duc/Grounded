<template>
  <div class="container mt-5">
    <h2 class="text-center mb-4">Create Item</h2>

    <form @submit.prevent="createItem">
      
      <!-- Name -->
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input
          type="text"
          v-model="form.name"
          class="form-control"
          id="name"
          required
        />
      </div>

      <!-- Code -->
      <div class="mb-3">
        <label for="code" class="form-label">Code</label>
        <input
          type="text"
          v-model="form.code"
          class="form-control"
          id="code"
        />
      </div>

      <!-- Upload File -->
      <div class="mb-3">
        <label class="form-label">Upload Image</label>
        <input type="file" class="form-control" @change="handleFileUpload" />
      </div>

      <!-- Item Type -->
      <div class="mb-3">
        <label for="itemTypeId" class="form-label">Type of Item</label>
        <select v-model="form.itemTypeId" class="form-select" required>
          <option disabled value="">Chọn loại item</option>
          <option
            v-for="type in itemTypes"
            :key="type.id"
            :value="type.id"
          >
            {{ type.name }}
          </option>
        </select>
      </div>

      <button type="submit" class="btn btn-primary w-100">Tạo mới</button>
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "CreateItem",

  data() {
    return {
      form: {
        name: "",
        code: "",
        itemTypeId: null, // ✔ sửa: dùng null, không dùng ""
      },
      file: null,
      itemTypes: [],
    };
  },

  created() {
    this.fetchItemTypes();
  },

  methods: {
    // Lấy danh sách item types
    fetchItemTypes() {
      axios
        .get("http://localhost:9999/items-type/")
        .then((res) => {
          this.itemTypes = res.data.result || res.data;
        })
        .catch((err) => {
          console.error("Lỗi lấy loại item:", err);
        });
    },

    // Nhận file upload
    handleFileUpload(event) {
      this.file = event.target.files[0];
    },

    // Submit form
    async createItem() {
      try {
        const formData = new FormData();

        formData.append("name", this.form.name);
        formData.append("code", this.form.code);
        formData.append("itemTypeId", Number(this.form.itemTypeId)); // ✔ ép số
        if (this.file) {
          formData.append("file", this.file);
        }

        // Debug xem FE gửi gì
        for (let p of formData.entries()) {
          console.log(p[0] + ":", p[1]);
        }

        await axios.post("http://localhost:9999/items/", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });

        alert("Tạo item thành công!");
        this.$router.push("/");
      } catch (err) {
        console.error("Lỗi tạo item:", err);
        alert("Tạo item thất bại.");
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
