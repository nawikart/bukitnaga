<template>
    <div>
        <v-parallax class="overlayed" src="/assets/images/slide1.jpg" height="250">
            <div class="middle text-xs-center px-3">
                <p class="font-weight-thin display-1">MAKE A RESERVATION</p>
                <p class="font-weight-thin subheading">For media enquiries, please contact: eMail: info@villabukitnaga.com, Tel - Bali // +62 811 388 11 22</p>
            </div>        
        </v-parallax>

        <v-layout row wrap font-weight-thin style="max-width:1200px;margin:auto;">
            <v-flex text-md-right text-xs-center xs12 md5 py-5 px-4 font-weight-thin subheading>
                <p class="font-weight-thin title blue--text mt-5">MAKE A RESERVATION</p>  
                <p>Villa Bukit Naga<br>
                Gianyar, Bali - Indonesia</p>

                <p>Villa Telephone: + 62 361 941118<br>
                Reservation Telephone: +62 811 388 1122<br>
                Email: info@villabukitnaga.com</p>

                <p>Time Zone GMT +8</p>                              
            </v-flex>            
            <v-flex xs12 md7 py-5 px-4>
                <!-- <p class="font-weight-thin subheading blue--text">please fill form below to send us enquiry.</p>   -->
                <v-form
                    ref="form"
                    v-model="valid"
                    lazy-validation
                >
                    <v-text-field
                    v-model="name"
                    :counter="10"
                    :rules="nameRules"
                    label="Name"
                    required
                    ></v-text-field>

                    <v-text-field
                    v-model="email"
                    :rules="emailRules"
                    label="E-mail"
                    required
                    ></v-text-field>

                    <v-select
                    v-model="select"
                    :items="items"
                    :rules="[v => !!v || 'Item is required']"
                    label="Item"
                    required
                    ></v-select>

                    <v-checkbox
                    v-model="checkbox"
                    :rules="[v => !!v || 'You must agree to continue!']"
                    label="Do you agree?"
                    required
                    ></v-checkbox>

                    <v-btn
                    color="blue"
                    dark
                    @click="validate"
                    >
                    Send
                    </v-btn>
                    <v-btn
                    color="warning"
                    @click="resetValidation"
                    >
                    Reset Validation
                    </v-btn>                    
                </v-form>
            </v-flex>
        </v-layout>            
    </div>
</template>

<script>
  export default {
    data: () => ({
      valid: true,
      name: '',
      nameRules: [
        v => !!v || 'Name is required',
        v => (v && v.length <= 10) || 'Name must be less than 10 characters'
      ],
      email: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+/.test(v) || 'E-mail must be valid'
      ],
      select: null,
      items: [
        'Item 1',
        'Item 2',
        'Item 3',
        'Item 4'
      ],
      checkbox: false
    }),

    methods: {
      validate () {
        if (this.$refs.form.validate()) {
          this.snackbar = true
        }
      },
      reset () {
        this.$refs.form.reset()
      },
      resetValidation () {
        this.$refs.form.resetValidation()
      }
    }
  }
</script>