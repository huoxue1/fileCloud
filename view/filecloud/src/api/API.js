const axios = require("axios");
module.exports = {
    base: process.env.VUE_APP_API,
    // 根据父节点获取所有目录下所有文件和文件夹
    getFiles: async function (parentID) {
        let response = await axios.get(this.base + "/files?parent_id=" + parentID)
        return response.data
    },

    // 获取根节点
    getRoot:async function(){
        return (await axios.get(this.base + "/root")).data
    },

    // 创建文件夹
    createDir: async function(name,parentID){
        return (await axios.post(this.base + "/mkdir", {"name": name, "parent_id": parentID})).data
    },

    // 获取文件或者文件夹信息
    getFile: async function(id){
        return (await axios.get(this.base + "/file?id=" + id)).data
    },

    deleteFile:async function(id){
        return (await axios.delete(this.base + "/delete_file?id=" + id)).data
    }
}