<template>
    <!-- INPUT PENYAKIT VIEW -->
    <v-app id="hasilprediksi" :style="{ background: $vuetify.theme.themes.dark.background }">
        <!-- NAVBAR -->
        <TheNavbar />

        

        <!-- FOOTER -->
        <TheFooter />
    </v-app>
</template>

<script>
  import TheNavbar from '../components/TheNavbar'
  import TheFooter from '../components/TheFooter'

  export default {
    name: 'HomeView',

    components: {
        TheNavbar,
        TheFooter,
    },

    methods: {
        scroll(refName) {
            const element = document.getElementById(refName);
            element.scrollIntoView({behavior: "smooth"});
        }
    }
  }
</script>

<style>

</style>