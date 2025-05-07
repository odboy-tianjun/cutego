<template>
  <div>
    <el-table
      :data="loginInfoList"
      height="250"
      border
      style="width: 100%">
      <el-table-column prop="ipAddr" label="登录IP地址" width="180"></el-table-column>
      <el-table-column prop="loginLocation" label="登录地点" width="180"></el-table-column>
      <el-table-column prop="browser" label="浏览器类型"></el-table-column>
      <el-table-column prop="os" label="操作系统"></el-table-column>
      <el-table-column prop="status" label="登录状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 1" type="danger">失败</el-tag>
          <el-tag v-if="scope.row.status === 0" type="success">成功</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="loginTime" label="登录时间"></el-table-column>
    </el-table>
    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="queryParams.pageNum"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />
  </div>
</template>

<script>
import {getLoginHistory} from '@/api/login'

export default {
  data() {
    return {
      loginInfoList: [],
      // 查询参数
      total: 0,
      queryParams: {
        pageNum: 1,
        pageSize: 10
      }
    }
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询用户登录记录列表 */
    getList() {
      getLoginHistory(this.queryParams).then(response => {
          this.loginInfoList = response.data.list
          this.total = response.data.total
        }
      )
    }
  }
}
</script>
