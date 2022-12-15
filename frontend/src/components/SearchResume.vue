<template>
  <div>
    <el-form ref="form" :inline="true" :model="form" class="demo-form-inline">
      <el-form-item label="关键字查询">
        <el-input v-model="form.searchContent" style="width: 200px"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form>
    <el-table
        :data="tableList"
        style="width: 100%">
      <el-table-column type="expand" v-if="isShow">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="">
              <span v-html="props.row.highlight"> </span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
          prop="_id"
          label="ID"
          width="220">
      </el-table-column>
      <el-table-column
          prop="_source.name"
          label="姓名"
          width="180">
      </el-table-column>
      <el-table-column
          prop="_source.school"
          label="学校">
      </el-table-column>
    </el-table>
  </div>

</template>

<script>
import api from '../utils/api'
export default {
  name: 'SearchResume',
  data() {
    return {
      form: {
        searchContent: '',
      },
      isShow:false,
      tableList:[],
      attachmentContentList: [],
    }
  },
  methods: {
    onSubmit() {
      const searchContent = this.form.searchContent
      this.isShow = searchContent !== '';
      console.log(this.isShow)
      this.$axios.get(api.CONTENTSEARCH.url + `?searchContent=${searchContent}&page=1&size=99`).then(res => {
        if (res.data.code === 0) {
          this.tableList = res.data.data.hits.hits;
          console.log( this.tableList)
        } else {
          this.$message.error(res.data.msg);
        }
      });
    },

  },
  created() {
    this.onSubmit();
  },
}
</script>

<style scoped>
</style>
