<template>

  <!-- HOME VIEW -->
  <v-app id="home" :style="{ background: $vuetify.theme.themes.dark.background }">

    <!-- NAVBAR -->
    <TheNavbar />

    <!-- CONTENT -->
    <v-container fluid>
      <v-row>

        <!-- HOME -->
        <v-col cols="6" class="home">
          <v-img src="1.png" contain max-height="600"></v-img>
        </v-col>
        <v-col cols="6">
          <h5 class="red--text text-darken-4 top">Selamat Datang Di</h5>
          <h1 class="white--text">DNA Pattern Matching</h1>
          <p class="grey--text">
            Program ini dibuat oleh Kelompok RucikaWavin
          </p>
          <v-btn tile color="#A7121D" @click="scroll('service')" dark>Get Started<v-icon bottom>mdi-arrow-down</v-icon></v-btn>
        </v-col>

        <!-- SERVICES -->
        <v-col cols="12" class="padd">
          <div class="service" id="service">
            <v-row>
              <v-col cols="12">

                <!-- Input Penyakit -->
                <div class="child child1">
                  <v-icon color="#A7121D" x-large class="ml-3">input</v-icon>
                  <h3 class="white--text ml-3 mt-4">Input Penyakit</h3>
                  <p class="grey--text ml-3 mt-6">
                    Input penyakit baru beserta sekuens DNA
                  </p>
                  <router-link class="routerLink" to="/inputpenyakit">
                    <v-btn color="#A7121D" depressed dark text>know more<v-icon right>mdi-arrow-right</v-icon></v-btn>
                  </router-link>
                </div>

                <!-- Tes DNA -->
                <div class="child child2">
                  <v-icon color="" x-large class="ml-3" dark>search</v-icon>
                  <h3 class="white--text ml-3 mt-4">Tes DNA</h3>
                  <p class="grey--text ml-3 mt-6">
                    Prediksi penyakit berdasarkan input DNA
                  </p>
                  <router-link class="routerLink" to="/tesdna">
                    <v-btn color="" depressed dark text>know more<v-icon right>mdi-arrow-right</v-icon></v-btn>
                  </router-link>
                </div>

                <!-- Hasil Prediksi -->
                <div class="child child1">
                  <v-icon color="#A7121D" x-large class="ml-3">article</v-icon>
                  <h3 class="white--text ml-3 mt-4">Hasil Prediksi</h3>
                  <p class="grey--text ml-3 mt-6">
                    Halaman yang menampilkan urutan hasil prediksi
                  </p>
                  <router-link class="routerLink" to="/hasilprediksi">
                    <v-btn color="#A7121D" depressed dark text>know more<v-icon right>mdi-arrow-right</v-icon></v-btn>
                  </router-link>
                </div>
              </v-col>
            </v-row>
          </div>
        </v-col>
      </v-row>
    </v-container>

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

<style scoped>
  /* CSS STYLE */
  .top {
    margin-top: 180px;
  }

  /* Home CSS */
  .home {
    width: 100%;
  }

  /* Service CSS */
  .service {
    width: 100%;
    height: 330px;
    background: linear-gradient(
      to right,
      #181818,
      #181818 50%,
      #111111 50%,
      #111111
    );
    text-align: center;
    padding: 2rem 2rem;
  }

  .child {
    display: inline-block;
    padding: 2rem 1rem;
    vertical-align: middle;
    text-align: left;
    margin-right: 8px;
  }
  .child1 {
    background-color: #1e1e1e;
    height: 250px;
    width: 250px;
  }
  .child2 {
    background-color: #ce1d2a;
    height: 250px;
    width: 250px;
  }

  .col-12.padd {
    padding: 12px 0 !important;
  }
  
</style>