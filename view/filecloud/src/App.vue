<template>
  <el-container style="height: 100%;margin: 0 auto">
    <el-header style="height: 15%;border: salmon solid 1px;padding-top: 1%">
      <el-row :gutter="20">
        <el-col :span="2"><span>fileCLoud</span></el-col>
        <el-col :span="2">
          <el-upload
              style="margin: 0 auto"
              :action="upload_action"
              :show-file-list="false"
              :auto-upload="true"
              :on-progress="uploadProgress"
              :on-success="()=>{flushFiles()}"
          >
            <el-button type="primary"
            >Upload File</el-button
            >
          </el-upload>
        </el-col>
        <el-col :span="2"><el-button type="primary" @click="createDir">Create Directory</el-button></el-col>
        <el-col :span="2"><el-button type="primary">Download File</el-button></el-col>
        <el-col :span="16"></el-col>

      </el-row>


      <el-row style="margin-top: 1.5%">
        <el-col span="24">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item v-for="p in path" :key="p.id">
              <span class="bre-item" @click="breadcrumbClick(p)" style="font-size: 30px">
                {{p.file_name}}
              </span>
            </el-breadcrumb-item>
          </el-breadcrumb>
        </el-col>
      </el-row>




    </el-header>
    <el-main style="height: 84.5%;border: seagreen solid 1px">
        <el-table :data="files" @row-dblclick="rowClick" style="width: 100%;">
              <el-table-column
                  label="Type"

              >
                <template #default="scope">

                      <el-icon v-if="scope.row.is_dir===1"><folder /></el-icon>
                      <el-icon v-else><document /></el-icon>
                </template>
              </el-table-column>
          <el-table-column prop="file_name" label="name" />

          <el-table-column label="File Size">
            <template #default="scope">
              {{get_size(scope.row.size)}}
            </template>
          </el-table-column>

          <el-table-column label="Upload Time">
            <template #default="scope">
              {{format_time(scope.row.upload_time)}}
            </template>
          </el-table-column>

          <el-table-column label="Action">
            <template #default="scope">
              <el-button size="small" :disabled="scope.row.is_dir===1" @click="downloadFile(scope.row.id)"
              >Download</el-button
              >
              <el-button
                  size="small"
                  type="danger"
                  @click="deleteFile(scope.row.id)"
              >Delete</el-button
              >
            </template>
          </el-table-column>
        </el-table>
    </el-main>
  </el-container>
</template>

<script>
import API from "./api/API"
import {Folder,Document} from "@element-plus/icons";

import {ElMessage, ElMessageBox, ElNotification} from "element-plus";
export default {
  components:{
    Folder,
    Document
  },
  data(){
    return{
      path:[],
      root:{},
      files:[],
      current_dir:{},

      searchValues:"",
      upload_action:API.base+"/upload?parent_id="
    }
  },
  watch:{
    current_dir(){
        this.upload_action = API.base+"/upload?parent_id="+this.current_dir.id
        this.freshPath()
        this.flushFiles()
    }
  },
  created() {
    API.getRoot().then((data)=>{
      if (data.ok){
        this.root = data.data
        this.path.push(data.data)
        this.current_dir = data.data
        API.getFiles(data.data.id).then(resp=>{
          if (resp.ok){
            this.files = resp.data
          }
        })
      }else {
        ElNotification.error({title:"error",message:data.msg})
      }
    })

  },
  methods:{

    // 删除文件
    deleteFile:async function (id) {
      ElMessageBox.confirm("Delete the File?","Warning",{confirmButtonText:"Ok",cancelButtonText:"Cancel",type:"warning"}).then(async () => {
        let resp = await API.deleteFile(id);
        if (resp.ok) {
              this.flushFiles()
        }else {
          ElNotification.error({message:resp.msg})
        }
      }).catch(()=>{

      })
    },

    // 下载文件
    downloadFile: function (id) {
        let element = document.createElement("a");
        element.href = API.base+"/download?id="+id
        element.click()
    },


    // 面包屑点击事件，跳转到对应文件夹
    breadcrumbClick:function (p) {
      this.current_dir = p
    },


    // table点击事件，如果是dir则进入对应文件夹
    rowClick:function (row, column, event) {
      console.log(row,column,event)
      if (row.is_dir === 1){
        this.current_dir = row
      }
    },

    // 刷新当前目录路径
    freshPath:async function () {
      this.path = []
      if (this.current_dir.id === this.root.id){
        this.path.push(this.root)
      }else {
        console.log(this.current_dir)
        let parent = this.current_dir.parent_id
        let resp;
        this.path.push(this.current_dir)
        while (parent !== ""){
          if (parent === undefined){
            break
          }
          resp = await API.getFile(parent)
          parent = resp.data.parent_id
          this.path.unshift(resp.data)
        }
      }
    },


    // 文件上传进度
    uploadProgress:function (event,file,fileList) {
      console.log(event)
      console.log(file)
      console.log(fileList)
    },

    // 创建目录
    createDir:function () {
        ElMessageBox.prompt("please input the directory name:","message",{confirmButtonText:"OK",cancelButtonText:"Cancel"}).then(value=>{
          API.createDir(value.value,this.current_dir.id).then((resp)=>{
            if (resp.ok){
              ElMessage.info("create a directory successful")
              this.flushFiles()
            }else {
              ElMessage.error(resp.msg)
            }
          })
        }).catch(()=>{

        })
    },

    // 刷新当前目录文件
    flushFiles:function () {
      API.getFiles(this.current_dir.id).then(resp=>{
        if (resp.ok){
          this.files = resp.data
        }
      })
    },
    // 根据时间戳格式化时间
    format_time: function formatTime(value) {
      if(value) {
        let date = new Date(value * 1000)	// 时间戳为秒：10位数
        //let date = new Date(value)	// 时间戳为毫秒：13位数
        let year = date.getFullYear()
        let month = date.getMonth() + 1 < 10 ? `0${date.getMonth() + 1}` : date.getMonth() + 1
        let day = date.getDate() < 10 ? `0${date.getDate()}` : date.getDate()
        let hour = date.getHours() < 10 ? `0${date.getHours()}` : date.getHours()
        let minute = date.getMinutes() < 10 ? `0${date.getMinutes()}` : date.getMinutes()
        let second = date.getSeconds() < 10 ? `0${date.getSeconds()}` : date.getSeconds()
        return `${year}-${month}-${day} ${hour}:${minute}:${second}`
      } else {
        return ''
      }
    },

    // 获取文件大小，根据字节
    get_size:function getfilesize(size) {
      if (!size)
        return "";
      const num = 1024.00; //byte
      if (size < num)
        return size + "B";
      if (size < Math.pow(num, 2))
        return (size / num).toFixed(2) + "K"; //kb
      if (size < Math.pow(num, 3))
        return (size / Math.pow(num, 2)).toFixed(2) + "M"; //M
      if (size < Math.pow(num, 4))
        return (size / Math.pow(num, 3)).toFixed(2) + "G"; //G
      return (size / Math.pow(num, 4)).toFixed(2) + "T"; //T
    },
  }
}
</script>



<style scoped>

.bre-item:hover{
  text-decoration: underline;
  color: blueviolet;
}
</style>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0 auto;

}

html,
body {
  margin: 0 auto;
  padding: 0;
  height: 100%;
}

</style>
