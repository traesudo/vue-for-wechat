<template>
  <div style="width: 256px">
    <a-button type="primary" style="margin-bottom: 16px" @click="toggleCollapsed">
      <MenuUnfoldOutlined v-if="state.collapsed" />
      <MenuFoldOutlined v-else />
    </a-button>
    <a-menu
      v-model:openKeys="state.openKeys"
      v-model:selectedKeys="state.selectedKeys"
      mode="inline"
      theme="dark"
      :inline-collapsed="state.collapsed"
      :items="items"
      @click = "handleClick"
    >
    ></a-menu>
  </div>
  <div>
    {{ data.resultText }}
  </div>
</template>

<script setup>
import { reactive, watch, h } from 'vue';
import {
  PieChartOutlined,
  DesktopOutlined,
  InboxOutlined,
} from '@ant-design/icons-vue';
import {OperationProcess} from '../../wailsjs/go/main/App'


const state = reactive({
  collapsed: false,
  selectedKeys: ['1'],
  openKeys: ['sub1'],
  preOpenKeys: ['sub1'],
});

const items = reactive([
  {
    key: '1',
    icon: () => h(PieChartOutlined),
    label: '查看好友列表',
    title: '查看好友列表',
  },
  {
    key: '2',
    icon: () => h(DesktopOutlined),
    label: '获取群聊',
    title: '获取群聊',
  },
  {
    key: '3',
    icon: () => h(InboxOutlined),
    label: '登陆',
    title: '登陆',
  },
]);


const handleClick = (item) => {
  OperationProcess(item.key).then((result) => {
    console.log("key-------",item.key)
    data.resultText = result
  }).catch((err) => {
    console.log("处理错误")
  });

};
const data = reactive({
  selectKey: 0,
  resultText: "Please enter your name below 👇",
});

const toggleCollapsed = () => {
  state.collapsed = !state.collapsed;
  state.openKeys = state.collapsed ? [] : state.preOpenKeys;
};
</script>
