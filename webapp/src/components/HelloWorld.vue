<template>
  <div class="hello">
    <h2>Add a todo</h2>
    <p>
      <input v-model="todo.title">
    </p>
    <button @click="createTodo">Create</button>
    <br>
    <h2>Todo list</h2>
    <ul>
      <ol v-for="todo in todos" :key="todo.id">{{ todo.title }}</ol>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HelloWorld',
  data () {
    return {
      todo: {
        title: ''
      },
      todos: []
    }
  },
  methods: {
    createTodo () {
      let url = 'http://localhost:4564/api/v1/todos/'
      let payload = {
        title: this.todo.title,
        completed: false
      }
      let vm = this
      axios.post(url, payload).then(function (resp) {
        vm.getTodos()
      })
    },
    getTodos () {
      let url = 'http://localhost:4564/api/v1/todos/'
      let vm = this
      axios.get(url).then(function (resp) {
        vm.todos = resp.data.todos
      })
    }
  },
  created () {
    this.getTodos()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
