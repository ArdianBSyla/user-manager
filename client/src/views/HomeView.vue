<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from 'axios';
import {useRouter} from "vue-router";

const users = ref([]);
const page = ref(1);
const size = ref(10);
const loading = ref(false);
const searchQuery = ref('');

const router = useRouter()

const fetchUsers = async () => {
  try {
    loading.value = true;
    const response = await axios.get<any>(
        `http://localhost:3000/api/v1/users?page=${page.value}&size=${size.value}`
    );
    users.value = response.data.data;
    loading.value = true;
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


</style>