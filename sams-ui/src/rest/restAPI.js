import axios from 'axios'

const axiosClient = axios.create({})

// apiService.js (JavaScript)
const BASE_URL = 'http://192.168.3.32:8080' // 后端 API 根路径

// 通用的请求方法
async function fetchAPI(endpoint, method = 'GET', body = null) {
  const headers = {
    'Content-Type': 'application/json',
    // 'Authorization': `Bearer ${token}`, // 如果需要身份认证，可以传入token
  }

  const options = {
    method,
    headers,
  }

  // 如果是 POST 或 PUT 请求，并且传入了 body，则需要加上请求体
  if (body) {
    options.body = JSON.stringify(body)
  }

  try {
    const response = await axiosClient(`${BASE_URL}${endpoint}`, options)

    // 如果响应状态不是 2xx，抛出错误
    if (response.status !== 200) {
      throw new Error(`HTTP error! Status: ${response.status}`)
    }

    const data = response.data
    return data
  } catch (error) {
    console.error('API 请求出错:', error)
    throw error // 抛出错误供调用方处理
  }
}

// 定义标准的 REST API 方法（GET, POST, PUT, DELETE）

export const restAPI = {
  // 获取数据 (GET)
  get(endpoint) {
    return fetchAPI(endpoint, 'GET')
  },

  // 创建数据 (POST)
  post(endpoint, body) {
    return fetchAPI(endpoint, 'POST', body)
  },

  // 更新数据 (PUT)
  put(endpoint, body) {
    return fetchAPI(endpoint, 'PUT', body)
  },

  // 删除数据 (DELETE)
  delete(endpoint) {
    return fetchAPI(endpoint, 'DELETE')
  },
}

export default restAPI
