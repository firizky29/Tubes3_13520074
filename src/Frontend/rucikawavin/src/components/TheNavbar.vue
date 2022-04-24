<template>
    <v-app-bar app color="#111111" dark flat class="px-8">
        <v-btn text>
            <v-icon left>mdi-fast-forward</v-icon>RucikaWavin
        </v-btn>
        <v-spacer></v-spacer>
        <v-btn text @click="scroll('home')">Home</v-btn>
        <v-btn text @click="scroll('service')">Project</v-btn>
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

</style>