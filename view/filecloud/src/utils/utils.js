module.exports = {
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

    fileToBase64: function (file,callback) {
        let reader = new FileReader();
        // 传入一个参数对象即可得到基于该参数对象的文本内容
        reader.readAsDataURL(file);
        reader.onload = function (e) {
            // target.result 该属性表示目标对象的DataURL
            callback(e.target.result)
        };
    }

}