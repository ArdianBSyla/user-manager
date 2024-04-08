<template>
  <div class="user-form">
    <h2>{{ edit ? "Edit User" : "Create User" }}</h2>
    <form @submit.prevent="submitData">
      <label for="first_name">First Name:</label>
      <input type="text" id="first_name" v-model="user.first_name" required />

      <label for="last_name">Last Name:</label>
      <input type="text" id="last_name" v-model="user.last_name" required />

      <label for="email">Email:</label>
      <input type="email" id="email" v-model="user.email" required />

      <button type="submit">{{edit ? "Edit" : "Create"}}</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from 'vue'
import axios from 'axios'
import {useRoute, useRouter} from 'vue-router'

    const route = useRoute()
const router = useRouter();

const edit = ref( !!route.query.id  )
    const user = ref({
      first_name: '',
      last_name: '',
      email: '',
      company_id: 1,
    })


    const createUser = async () => {
      try {
        const response = await axios.post('http://localhost:3000/api/v1/users', user.value)
        console.log(response.data)
        router.push('/')
      } catch (error) {
        console.error('Error creating user:', error)
      }
    }

    const editUser = async (id: number) => {
      try {
        const response = await axios.put('http://localhost:3000/api/v1/users/' + id, user.value)
        console.log(response.data)
        router.push('/')
      } catch (error) {
        console.error('Error creating user:', error)
      }
    }

    const submitData = async () => {
  const id = route.query.id;
      if(id) {
        //Call an endpoint to edit
        editUser(Number(id))
      } else {
        createUser()
      }
    }

    onMounted(() => {
      if(route.query.id) {
        axios.get(`http://localhost:3000/api/v1/users/${route.query.id}`)
          .then(response => {
            console.log(response.data.data)
            user.value = response.data.data
          })
          .catch(error => {
            console.error(`Error fetching user: ${error}`)
          })
      }
    })

</script>

<style scoped>
.user-form {
  max-width: 400px;
  margin: 0 auto;
}

.user-form h2 {
  font-size: 1.5rem;
  margin-bottom: 20px;
}

.user-form label {
  display: block;
  margin-bottom: 5px;
}

.user-form input[type="text"],
.user-form input[type="email"] {
  width: 100%;
  padding: 8px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.user-form button {
  background-color: #4CAF50;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.user-form button:hover {
  background-color: #45a049;
}
</style>
