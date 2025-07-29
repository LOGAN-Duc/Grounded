<template>
  <nav class="menu">
    <ul>
      <li><router-link to="/">Danh Sách Items</router-link></li>
      <li><router-link to="/resources">Danh Sach Resources</router-link></li>
      <li><router-link to="/about">Giới Thiệu</router-link></li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'NavMenu', // Đổi tên thành NavMenu
};
</script>

<style scoped>
.menu {
  background-color: #f8f9fa;
  padding: 10px;
}
.menu ul {
  list-style-type: none;
  padding: 0;
}
.menu li {
  margin: 5px 0;
}
.menu a {
  text-decoration: none;
  color: #007bff;
}
.menu a:hover {
  text-decoration: underline;
}
</style>