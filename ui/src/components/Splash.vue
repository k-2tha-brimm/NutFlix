<template>
  <div class="hello">
    <h1>BEHOLD!</h1>
    <h2>MOVIES JSON</h2>
    <button @click="getMovies">Get Movies</button>
    <div class="movies" v-for="movie in movies" v-bind:key="movie.id">
      {{ movie.title }}
    </div>
  </div>
</template>

<script>
import {HTTP} from '../http-constants';
export default {
  name: 'Splash',
  data () {
    return {
      movies: [],
      errors: ''
    }
  },
  methods: {
    getMovies: function () {
      HTTP.get('/api/movies')
      .then(res => {
        this.movies = res.data
      })
      .catch(e => {
        this.errors = e
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
