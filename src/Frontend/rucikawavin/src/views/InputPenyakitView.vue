<template>
    <!-- INPUT PENYAKIT VIEW -->
    <v-app id="inputpenyakit" :style="{ background: $vuetify.theme.themes.dark.background }">
        <!-- NAVBAR -->
        <TheNavbar />

        <v-container fluid>
            <v-row>
                <v-col cols="12" class="padd">
                    <div class="formContainer" id="formContainer">
                        <form class="col-md-8 rounded px-5 py-4 shadow bg-white form" @submit="onSubmit">
                            <h1 class="white--text">Tambah Penyakit</h1>
                            <div class="col-12 form-group">
                                <label class="col-form-label col-form-label-lg white--text">Nama Penyakit <span class="text-danger">*</span></label>
                                <input @focus="errorName = false" type="text" class="form-control form-control-lg white--text" :class="{'error-border': errorName}" placeholder="Nama Penyakit" v-model="namapenyakit" />
                            </div>
                            <div class="col-12 form-group">
                                <label class="col-form-label col-form-label-lg white--text">Sequence DNA <span class="text-danger">*</span></label>
                                <input @focus="errorContent = false" accept=".txt" type="file" class="form-control form-control-lg white--text" color="#A7121D" :class="{'error-border': errorContent}" placeholder="Sequence DNA" name="sequencedna" @change="onFilePicked">
                            </div>
                            <div class="col-12 form-group text-center">
                                <v-btn tile color="#A7121D" type="submit" dark>Submit<v-icon right>upload</v-icon></v-btn>
                            </div>
                        </form>
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
        name: 'InputPenyakitView',
        data() {
            return {
                namapenyakit:"",
                file: null,
                content:"",
                errorName:false,
                errorContent:false,
            }
        },

        components: {
            TheNavbar,
            TheFooter,
        },

        methods: {
            async onSubmit(event) {
                event.preventDefault();
                this.errorName = false
                this.errorContent = false
                if (!this.namapenyakit || !this.content || !this.file) {
                    if(!this.namapenyakit){
                        this.errorName = true
                    } else{
                        this.errorContent = true
                    }
                    await this.$alert("All fields must not empty.", "Error", "error");
                } else {
                  try{
                    await axios.post('/inputpenyakit', {
                      DiseaseName: this.namapenyakit,
                      DNA: this.content
                    })
                    await this.$alert("Disease data has been added successfully", "Success", "success");
                  }catch (e) {
                      console.log(e.response.data)
                  }
                }
            },
            onFilePicked(event) {
                this.errorContent = false
                const files = event.target.files
                this.file = files[0]
                const reader = new FileReader();
                if (this.file.name.includes(".txt")) {
                  reader.onload = (res) => {
                    this.content = res.target.result;
                  };
                  reader.onerror = (err) => console.log(err);
                  reader.readAsText(this.file);
                } else {
                   this.errorContent = true
                }
            }
        }
    }
</script>

<style>
    /* Form CSS */
    .formContainer {
        width: 100%;
        height: 650px;
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

    input[type="text"], input[type="file"]  {
        all: unset;
        width: 95%;
        height: 100%;
        border: 1px solid;
        margin-top: 1rem;
        margin-bottom: 1rem; 
        padding: 1rem;
    }

    input.error-border{
      border: 1px solid red;
    }

    input[type="file"]::file-selector-button {
        border: 2px solid #A7121D;
        background-color: #A7121D;
        transition: 1s;
        margin-right: 1rem;
    }

</style>