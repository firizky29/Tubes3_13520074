<template>
    <v-app-bar app color="#111111" dark flat class="px-8">
        <v-btn text>
            <v-icon left>mdi-fast-forward</v-icon>RucikaWavin
        </v-btn>
        <v-spacer></v-spacer>
        <router-link class="routerLink" to="/"><v-btn text>Home</v-btn></router-link>
        <router-link class="routerLink" to="/inputpenyakit"><v-btn text>Input Penyakit</v-btn></router-link>
        <router-link class="routerLink" to="/tesdna"><v-btn text>Tes DNA</v-btn></router-link>
        <router-link class="routerLink" to="/hasilprediksi"><v-btn text>Hasil Prediksi</v-btn></router-link>
    </v-app-bar>
</template>

<script>
export default {
    methods: {
        scroll(refName) {
            const element = document.getElementById(refName);
            element.scrollIntoView({behavior: "smooth"});
        }
    }
}
</script>

<style>
    .routerLink{
        text-decoration: none;
    }
</style>