// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';
import ItemList from '../components/item/ItemList.vue'
import CreateItem from '../components/item/CreateItem.vue'
import EditItem from '../components/item/EditItem.vue'

import ResourceList from '../components/resource/ResourceList.vue'
import ResourceCreate from '../components/resource/ResourceCreate.vue'
import ResourceEdit from '../components/resource/ResourceEdit.vue'

import CreateItemResouce from '../components/ItemResource/Create.vue'


const routes = [
  { path: '/', component: ItemList },
  { path: '/items/create', component: CreateItem },
  { path: '/items/edit/:id', component: EditItem , props : true},
  { path: '/resources', component: ResourceList },
  { path: '/resources/create', component: ResourceCreate },
  { path: '/resources/edit/:id', component: ResourceEdit, props : true },
  { path:'/itemresource/create/:id', component: CreateItemResouce}
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
