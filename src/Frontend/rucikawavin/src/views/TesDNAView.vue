<template>
    <!-- TES DNA VIEW -->
    <v-app id="tesdna" :style="{ background: $vuetify.theme.themes.dark.background }">
        <!-- NAVBAR -->
        <TheNavbar />

        <v-container fluid>
            <v-row>
                <v-col cols="12" class="padd">
                    <div class="formContainer" id="formContainer">
                        <form class="col-md-8 rounded px-5 py-4 shadow bg-white form" @submit="onSubmit">
                            <h1 class="white--text">Tes DNA</h1>

                            <!-- Nama Pengguna -->
                            <div class="col-12 form-group">
                                <label class="col-form-label col-form-label-lg white--text">Nama Pengguna <span class="text-danger">*</span></label>
                                <input type="text" class="form-control form-control-lg white--text" placeholder="Nama Pengguna" v-model="namapengguna" />
                            </div>

                            <!-- FIle -->
                            <div class="col-12 form-group">
                                <label class="col-form-label col-form-label-lg white--text">Sequence DNA <span class="text-danger">*</span></label>
                                <input accept=".txt" type="file" class="form-control form-control-lg white--text" color="#A7121D" placeholder="Sequence DNA" name="sequencedna" @change="onFilePicked">
                            </div>

                            <!-- Penyakit -->
                            <div class="col-12 form-group">
                                <label class="col-form-label col-form-label-lg white--text">Penyakit <span class="text-danger">*</span></label>
                                <input type="text" class="form-control form-control-lg white--text" color="#A7121D" placeholder="Penyakit" name="penyakit" v-model="penyakit">
                            </div>

                            <!-- Algoritma -->
                            <div class="col-12">
                                <label class="col-form-label col-form-label-lg white--text">Algoritma <span class="text-danger">*</span></label> <br />
                                
                                <div class="radioContainer">
                                    <input type="radio" name="algoritma" v-model="algoritma" value="KMP">
                                    <label class="white--text">KMP</label>
                                </div>
                                <div class="radioContainer">
                                    <input type="radio" name="algoritma" v-model="algoritma" value="Boyer Moore">
                                    <label class="white--text">Boyer Moore</label>
                                </div>
                            </div>
                            <div class="col-12 form-group text-center">
                                <v-btn tile color="#A7121D" type="submit" dark>Submit<v-icon right>upload</v-icon></v-btn>
                            </div>
                            <h1 class="white--text" v-if="!resultHidden">Hasil Tes</h1>
                            <div class="col-12 form-group" v-if="!resultHidden">
                                <input type="text" class="form-control form-control-lg white--text" color="#A7121D" name="hasiltes" :value="hasiltes" readonly>
                            </div>
                        </form>"
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
        name: 'TesDNAView',
        data() {
            return {
                namapengguna:"",
                dnapengguna:"",
                penyakit: "",
                hasiltes: "",
                file: null,
                resultHidden: true,
                algoritma: ""
            }
        },

        components: {
            TheNavbar,
            TheFooter,
        },

        methods: {
            async onSubmit(e) {
                e.preventDefault();
                if (!this.namapengguna || !this.dnapengguna || !this.penyakit || !this.algoritma) {
                    await this.$alert("Please make sure your input is not empty.", "Error", "error");
                } else {
                  this.resultHidden = false;
                  try{
                      const data = await axios.post('/tesdna', {
                          UserName: this.namapengguna,
                          DNA: this.dnapengguna,
                          DiseasePrediction: this.penyakit,
                          Algorithm: this.algoritma
                      })
                      this.hasiltes = data.data.result
                  }catch (e) {
                      await this.$alert(e.response.data.error, "Error", "error")
                      this.resultHidden = true;
                  }
                }
            },
            onFilePicked(event) {
                const files = event.target.files
                this.file = files[0]
                const reader = new FileReader();
                if (this.file.name.includes(".txt")) {
                  reader.onload = (res) => {
                    this.dnapengguna = res.target.result;
                  };
                  reader.onerror = (err) => console.log(err);
                  reader.readAsText(this.file);
                }
            }
        }
    }
</script>

<style>
    /* Form CSS */
    .formContainer {
        width: 100%;
        height: 100%;
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

    .form {
        display: inline-block;
        background-color: #1e1e1e;
        padding: 2rem 2rem;
        vertical-align: middle;
        text-align: left;
    }

    .col-12.padd {
        padding: 12px 0 !important;
    }

    input[type="text"], input[type="file"] {
        all: unset;
        width: 95%;
        height: 100%;
        border: 1px solid;
        margin-top: 1rem;
        margin-bottom: 1rem; 
        padding: 1rem;
    }

    input[type="file"]::file-selector-button {
        border: 2px solid #A7121D;
        background-color: #A7121D;
        transition: 1s;
        margin-right: 1rem;
    }

    .radioContainer {
        margin: 1rem 1rem 1rem 0rem;
    }

    .radioContainer input {
        margin-right: 1rem;
    }
</style>