<template>
    <!-- INPUT PENYAKIT VIEW -->
    <v-app id="hasilprediksi" :style="{ background: $vuetify.theme.themes.dark.background }">
        <!-- NAVBAR -->
        <TheNavbar />
        
        <!-- CONTENT -->
        <v-container fluid>
            <v-row class="hasilprediksi">
        
                <!-- SEARCHBAR -->
                <v-col cols="12">
                    <div class="searchbar" id="searchbar">
                        <h1 class="mb-10 white--text">Pencarian Hasil Prediksi</h1>
                        <form class="col-md-8 rounded px-5 py-4 shadow bg-white form" @submit="onSubmit">
                            <div cols="6" class="form-group">
                                <input type="text" class="form-control form-control-lg white--text" placeholder="&quot;dd-mm-yyyy <nama penyakit>&quot;, &quot;<nama penyakit>&quot;, &quot;dd-mm-yyyy&quot;" v-model="search" />
                            </div>
                            <div cols="6" class="form-group text-center">
                                <v-btn tile color="#A7121D" type="submit" dark>Search <v-icon right>search</v-icon></v-btn>
                            </div>
                        </form>
                    </div>
                </v-col>

                <!-- SEARCH RESULT -->
                <v-col cols="12" class="resultcontainer">
                    <div class="searchresult" v-for="result of searchResults" :key="result.hasil">
                        <div class="col-md-8 rounded px-5 py-4 shadow bg-white result white--text">
                            {{result.hasil}}
                        </div>
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
    import axios from "axios";

    export default {
        name: 'HasilPrediksiView',

        data() {
            return {
                results: [
                    // // dummy
                    {hasil: 'ayam'},
                    {hasil: 'babi'},
                    {hasil: 'anjing'},
                    {hasil: 'sapi'},
                    {hasil: 'bola'},
                    //
                    // // {id: 1, tanggal: '10 April 2022', pasien: 'Fulan', penyakit: 'HIV', hasil: 'True'},
                    // // {id: 2, tanggal: '11 April 2022', pasien: 'Kamal', penyakit: 'Kanker', hasil: 'False'},
                    // // {id: 3, tanggal: '12 April 2022', pasien: 'Entah', penyakit: 'Serangan Jantung', hasil: 'False'},
                    // // {id: 4, tanggal: '13 April 2022', pasien: 'Jamal', penyakit: 'Impotensi', hasil: 'True'},
                    // // {id: 5, tanggal: '14 April 2022', pasien: 'Yubai', penyakit: 'Gangguan Kehamilan', hasil: 'True'},
                    // // {id: 6, tanggal: '15 April 2022', pasien: 'Hika', penyakit: 'Gangguan Janin', hasil: 'False'},
                ],
                searchResults: [],
                search: ""
            }
        },

        components: {
            TheNavbar,
            TheFooter,
        },

        methods: {
            async onSubmit(e) {
                e.preventDefault();
                try{
                    if (this.search !== "") {
                        // Jika input tidak kosong
                        const data = await axios.post('/hasilprediksi', {
                            Query: this.search 
                        })

                        // Dimasukkan ke dalam array searchResults
                        this.searchResults = data.data.Result
                    } else {
                        // Jika input kosong, masukan seluruh data ke dalam array searchResults
                        this.searchResults = this.results
                    }
                }catch (e) {
                    await this.$alert(e.response.data.error, "Error", "error")
                }
            },
        }
    }
</script>

<style>
    /* Searchbar CSS */
    .hasilprediksi {
        width: 100%;
        min-height: 650px;
        background: linear-gradient(
        to bottom,
            #181818,
            #181818 50%,
            #111111 50%,
            #111111
        );
        text-align: center;
        padding: 2rem;
    }

    .searchbar {
        width: 100%;
        text-align: center;
        padding: 2rem;
    }

    .form {
        display: inline-block;
        background-color: #1e1e1e;
        padding: 2rem 2rem;
        vertical-align: middle;
        text-align: left;
    }

    .resultcontainer {
        height: 100%;
        text-align: center;
        padding: 2rem;
    }

    .result {
        display: inline-block;
        background-color: #1e1e1e;
        padding: 2rem 2rem;
        margin-bottom: 1rem;
        vertical-align: middle;
        text-align: left;
    }

    input[type="text"] {
        all: unset;
        width: 95%;
        height: 100%;
        border: 1px solid;
        margin-top: 1rem;
        margin-bottom: 1rem; 
        padding: 1rem;
    }
</style>