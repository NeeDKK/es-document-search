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
      <el-table-column type="expand" v-if="isShowSearch">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="">
              <span v-html="props.row.highlight"> </span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column type="expand" v-if="isShowAll">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="">
              <span v-html="props.row._source.attachment.content"> </span>
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
    <div class="block" style="text-align: right;margin-top: 30px">
      <el-pagination
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          background
          layout="prev, pager, next"
          :page-size="pagesize"
          :total="total">
      </el-pagination>
    </div>
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
      isShowAll: false,
      isShowSearch: false,
      tableList: [],
      attachmentContentList: [],
      total: 0,
      pagesize:20,
      currentPage: 1,
    }
  },
  methods: {
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage
      this.onSubmit()
    },
    onSubmit() {
      const searchContent = this.form.searchContent
      this.$axios.get(api.CONTENTSEARCH.url + `?searchContent=${searchContent}&page=` + this.currentPage + `&size=20`).then(res => {
        if (res.data.code === 0) {
          this.tableList = res.data.data.hits.hits;
          this.total = res.data.data.hits.total.value;
          if (searchContent === '') {
            this.isShowAll = true
            this.isShowSearch = false
            for (let i = 0; i < this.tableList.length; i++) {
              let split = this.tableList[i]._source.attachment.content.split("\n");
              let contents = ''
              for (let j = 0; j < split.length; j++) {
                if (!split[j].match(/^[ ]*$/)) {
                  contents += split[j] + "</br>"
                }
              }
              this.tableList[i]._source.attachment.content = contents
            }
          } else {
            this.isShowAll = false
            this.isShowSearch = true
          }
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
