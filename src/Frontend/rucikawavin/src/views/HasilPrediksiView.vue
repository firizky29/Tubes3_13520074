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
                        <div class="col-md-8 rounded px-5 py-4 shadow bg-white form">
                            <div class="col-12 form-group">
                                <input type="text" class="form-control form-control-lg white--text" placeholder="&quot;dd-mm-yyyy nama_penyakit&quot;, &quot;nama_penyakit&quot;, &quot;dd-mm-yyyy&quot;" v-model="search" />
                            </div>
                        </div>
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

    export default {
        name: 'HomeView',

        data() {
            return {
                results: [
                    // dummy
                    {hasil: 'ayam'},
                    {hasil: 'babi'},
                    {hasil: 'anjing'},
                    {hasil: 'sapi'},
                    {hasil: 'bola'},

                    // {id: 1, tanggal: '10 April 2022', pasien: 'Fulan', penyakit: 'HIV', hasil: 'True'},
                    // {id: 2, tanggal: '11 April 2022', pasien: 'Kamal', penyakit: 'Kanker', hasil: 'False'},
                    // {id: 3, tanggal: '12 April 2022', pasien: 'Entah', penyakit: 'Serangan Jantung', hasil: 'False'},
                    // {id: 4, tanggal: '13 April 2022', pasien: 'Jamal', penyakit: 'Impotensi', hasil: 'True'},
                    // {id: 5, tanggal: '14 April 2022', pasien: 'Yubai', penyakit: 'Gangguan Kehamilan', hasil: 'True'},
                    // {id: 6, tanggal: '15 April 2022', pasien: 'Hika', penyakit: 'Gangguan Janin', hasil: 'False'},
                ],
                search: ""
            }
        },

        components: {
            TheNavbar,
            TheFooter,
        },

        computed: {
            searchResults() {
                let res = []
                if (this.search !== "") {
                    // contoh program
                    res = this.results.filter(p =>
                        p.hasil.toLowerCase().includes(this.search.toLowerCase())
                        // p.tanggal.toLowerCase().includes(this.search.toLowerCase()) ||
                        // p.penyakit.toLowerCase().includes(this.search.toLowerCase())
                    )
                } else {
                    res = this.results
                } 
                return res
            }
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