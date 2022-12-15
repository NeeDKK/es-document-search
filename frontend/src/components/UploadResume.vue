<template>
  <div>
    <el-upload
        class="upload-demo"
        drag
        action="#"
        :http-request="uploadHttpRequest"
        multiple>
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <div class="el-upload__tip" slot="tip">只能上传word/pdf文件</div>
    </el-upload>
  </div>
</template>

<script>
export default {
  name: 'UploadResume',
  data() {
    return {
      customFile: '',
    }
  },
  methods:{
    uploadHttpRequest(param) {
      this.customFile = param.file
      const data = new FormData()
      const fileUps = this.customFile
      data.append('file', fileUps)
      this.$axios.postForm('/uploadFile', data).then(res => {
        if (res.data.code === 0) {
          this.$message.success('上传成功')
          this.customFile=''
        } else {
          this.$message.error(res.data.msg)
          this.customFile=''
        }
      })
    },
  }
}
</script>

<style scoped>
</style>
