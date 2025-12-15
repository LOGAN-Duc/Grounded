import { createRouter, createWebHistory } from 'vue-router';
import ItemList from '../components/item/ItemList.vue';
import CreateItem from '../components/item/CreateItem.vue';
import EditItem from '../components/item/EditItem.vue';

import ResourceList from '../components/resource/ResourceList.vue';
import ResourceCreate from '../components/resource/ResourceCreate.vue';
import ResourceEdit from '../components/resource/ResourceEdit.vue';

import CreateItemResource from '../components/ItemResource/Create.vue';
import View from '@/components/landMark/view.vue';
import CreateLandmark from '@/components/landMark/create.vue';
import CreateLandMarkItem from '@/components/landMarkItem/createLandMarkItem.vue';
import LandmarkView from '@/components/landMark/LandmarkView.vue'; 

const routes = [
  { path: '/', component: ItemList },
  { path: '/items/create', component: CreateItem },
  { path: '/items/edit/:id', component: EditItem , props: true },
  { path: '/resources', component: ResourceList },
  { path: '/resources/create', component: ResourceCreate },
  { path: '/resources/edit/:id', component: ResourceEdit, props: true },
  { path: '/itemresource/create/:id', component: CreateItemResource },
  { path: '/landmarks', component: View },
  { path: '/landmarks/create', component: CreateLandmark },
  { path: '/landmarks/:id/add-items-resources', component: CreateLandMarkItem, props: true },
   { path: '/landmarks/view/:id', component: LandmarkView, props: true },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
