<template>
  <div class="master">
    <div class="avatar-container">
      <el-avatar class="avatar">
        <img
          src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png"
        />
      </el-avatar>
    </div>
    <div class="content">
      <!-- 其他内容 -->
      Content here
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
.master {
  height: 12vh; /* 使用视口高度的20% */
  width: 100%;
  background-color: #f5f5f5;
  border-bottom: 1px solid #e5e5e5;
  display: flex;
  align-items: center; /* 垂直居中 */
}

.avatar-container {
  width: 12vh; /* 使用父容器宽度的10% */
  display: flex;
  align-items: center; /* 垂直居中 */
}

.avatar {
  width: 100%;
  height: auto; /* 高度自动 */
}

.avatar img {
  width: 100%;
  height: auto; /* 确保图片按比例填充容器 */
  object-fit: cover; /* 确保图片按比例填充容器 */
}

.content {
  flex-grow: 1; /* 使内容占据剩余空间 */
  display: flex;
  align-items: center; /* 垂直居中对齐 */
  padding-left: 10px; /* 为内容区域添加左边距 */
}
</style>
