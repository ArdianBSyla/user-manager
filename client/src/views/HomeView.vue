<script setup lang="ts">
import {ref, onMounted, watch} from 'vue';
import axios from 'axios';
import {useRouter} from "vue-router";

const users = ref([]);
const page = ref(1);
const size = ref(10);
const total = ref(0);
const loading = ref(false);
const searchQuery = ref('');

const router = useRouter()

watch(searchQuery, () => {
  debouncedSearch();
});

watch(page, () => {
  fetchUsers()
})

watch(size, () =>{
  fetchUsers()
})

const fetchUsers = async () => {
  try {
    loading.value = true;
    const response = await axios.get<any>(
        `http://localhost:3000/api/v1/users?page=${page.value}&size=${size.value}&query=${searchQuery.value}`
    );
    users.value = response.data.data;
    loading.value = false;
    total.value = response.data.total;
    console.log(response.data.data);
  } catch (error) {
    console.error('Error fetching users:', error);
  }
};

const deleteUser = async (user: any) => {
  try {
    const response = await axios.delete(`http://localhost:3000/api/v1/users/${user.id}`);
    fetchUsers();
  } catch (error) {
    console.error('Failed to delete user', error);
  }
};

const editUser = (user: any) => {
  //TODO: navigate to edit user view
  router.push('/edit?id=' + user.id);
};

const debounce = (func: Function, delay: number) => {
  let timeoutId: ReturnType<typeof setTimeout>;
  return (...args: any[]) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
      func(...args);
    }, delay);
  };
};

const debouncedSearch = debounce(fetchUsers, 500);

const handleSizeChange = (currentPage: number, pageSize: number) => {
  size.value = pageSize;
}

const handleCurrentChange = (currentPage: number) => {
  page.value = currentPage
}

onMounted(fetchUsers);
</script>


<template>
  <div class="users">
    <div class="header">
      <h1>Users</h1>
      <input type="text" v-model="searchQuery" placeholder="Search..." class="search-input">
      <RouterLink to="/edit">
        <button class="add-user-btn">Add User</button>
      </RouterLink>

    </div>
    <div class="loader">Loading ...</div>
    <el-table
        :data="users"
        width="80vw"
        style="width: 80vw">
      <el-table-column
          prop="first_name"
          label="First Name"
          width="180">
      </el-table-column>
      <el-table-column
          prop="last_name"
          label="Last Name"
          width="180">
      </el-table-column>
      <el-table-column
          prop="email"
          label="Email">
      </el-table-column>
      <el-table-column
          label="Actions">
        <template #default="scope">
          <el-button type="primary" size="small" @click="editUser(scope.row)">Edit</el-button>
          <el-button type="danger" size="small" @click="deleteUser(scope.row)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="page"
        :page-sizes="[10, 20, 30, 50]"
        :page-size="size"
        layout="total, sizes, prev, pager, next, jumper"
        class="pagination"
        background
        :total="total">
    </el-pagination>
  </div>

</template>



<style scoped>
.users {
  max-width: 800px;
  margin: 0 auto;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.add-user-btn {
  background-color: #4CAF50;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.add-user-btn:hover {
  background-color: #45a049;
}

.search-input {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  width: 200px;
}

.loader {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.pagination {
  margin-top: 10px;
}



</style>