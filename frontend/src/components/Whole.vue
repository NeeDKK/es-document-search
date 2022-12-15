<template>
  <div id="wrapper">
    <!-- 页面头部部分 -->
    <div class="header">
      <div class="logo">人才库</div>
      <!-- 水平一级菜单 -->
      <div style="float:left;">
        <el-menu
            text-color="#000000"
            active-text-color="#3989fa"
            :default-active="toIndex"
            @select="handleSelect">
          <el-menu-item v-for="(item, index) in itemList" :index="item.path" :key="index">
            <span slot="title">{{ item.title }}</span>
          </el-menu-item>
        </el-menu>
      </div>
    </div>
    <!-- 页面左侧二级菜单栏，和主体内容区域部分 -->
    <el-main>
      <router-view></router-view>
    </el-main>
  </div>
</template>

<script>

export default {
  data() {
    return {
      itemList: [    // 水平一级菜单栏的菜单
        {path: '/', title: '首页'},
        {path: '/search', title: '内容查询'},
        {path: '/upload', title: '上传简历'},
      ],
    }
  },
  computed: {
    toIndex() {  // 根据路径绑定到对应的一级菜单，防止页面刷新重新跳回第一个
      return '/' + this.$route.path.split('/')[1];
    },
  },
  methods: {
    handleSelect(path) {  // 切换菜单栏
      this.$router.push({
        path: path
      });
    },
  }
}
</script>

<style scoped>
#wrapper {
  width: 100%;
  height: 100%;
  background: #f0f0f0;
}

.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
}

.header .logo {
  text-align: left;
  margin-left: 60px;
  margin-top: 0px;
  height: 29px;
  width: 160px;
  vertical-align: middle;
}

.el-menu.el-menu {
  border-bottom: none !important;
  float: left;
  margin-left: 50px;
  background: transparent;
}

.el-menu--horizontal > .el-menu-item.is-active {
  /* border-bottom: 2px solid #3989fa;
  color: #3989fa; */
  font-weight: bold;
}

.el-menu--horizontal > .el-menu-item {
  font-size: 16px;
  margin: 0 15px;
}

</style>