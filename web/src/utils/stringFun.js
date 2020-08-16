import {formatTimeToStr} from "@/utils/data";

export const toUpperCase = (str) => {
    if (str[0]) {
        return str.replace(str[0], str[0].toUpperCase())
    } else {
        return ""
    }
}

// 驼峰转换下划线
export const toSQLLine = (str) => {
    if (str=="ID") return "ID"
    return str.replace(/([A-Z])/g,"_$1").toLowerCase();
  }

  // 下划线转换驼峰
  export const  toHump = (name) => {
    return name.replace(/\_(\w)/g, function(all, letter){
        return letter.toUpperCase();
    });
}

export const randomType = () => {
    let type = ['success', 'info', 'danger', 'warning', '']
    let seed = Math.floor(Math.random() * 5);
    return type[seed]
}

export const formatDate = (time) => {
    if (time != null && time !== "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd");
    } else {
        return "";
    }
}

export const formatDateTime = (time) => {
    if (time != null && time !== "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
    } else {
        return "";
    }
}

export  const formatOnline = (bool) => {
    if (bool != null) {
        return bool ? "在线" :"离线";
    } else {
        return "";
    }
}