<template>
  <div :class="{ 'nav-open': $sidebar.showSidebar } ">
    <img ref="img" :src="urlLogo" style="display:none" />
    <router-view v-if="loadContext"></router-view>
  </div>

</template>
<script>
  import ColorThief from 'colorthief'
  const colorThief = new ColorThief()
  export default {
    name: 'app',
    data() {
      return {
        projectName: "",
        urlLogo: "",
        loadContext: false,
        palette: [],
        skin: {
          background: "",
          card: "",
          panel: "",
          wallpaper: "",
          toolbar: "",
          text: "",
          dark: "",
          transparent: "",
          button: ""
        },
        takeRootCSS: ""
      }
    },
    async mounted() {
      var pathname = window.location.pathname;
      this.projectName = pathname.split("/")[2]
      this.urlLogo = `/api/project/avatar/${this.projectName}.png`;

      await this.axios.get(this.urlLogo).then(rs => {
        if (rs.data == 'EOF') {
          window.location.href = "/"
        }
        var avatar = new Image();
        avatar.onload = this.loadAvatarProject();
        avatar.src = this.urlLogo;
        var link = document.querySelector("link[rel*='icon']") || document.createElement('link');
        link.type = 'image/x-icon';
        link.rel = 'shortcut icon';
        link.href = this.urlLogo;
        document.getElementsByTagName('head')[0].appendChild(link);
        document.title = this.projectName + "@DaVinci";

      })
    },
    methods: {
      loadAvatarProject() {
        //console.log("Cargando Paleta de Colores [Still]");
        if (typeof this.$refs.img === "undefined") {
          setTimeout(() => {
            this.loadAvatarProject();
          }, 80)
        } else {
          //console.log("Cargando Paleta de Colores [OK]");
          try {

            this.palette = colorThief.getPalette(this.$refs.img);
            if (this.palette == null) {
              setTimeout(() => {
                this.loadAvatarProject();
              }, 80)
            } else {
              this.palette.sort();
              this.palette.reverse();
              for (var pal in this.palette) {
                //console.log(JSON.stringify(this.palette[pal]));
              }

              this.skin.card = `rgb(${this.palette[1][0]},${this.palette[1][1]},${this.palette[1][2]})`;
              this.skin.background = `rgb(${this.palette[2][0]},${this.palette[2][1]},${this.palette[2][2]})`;
              this.skin.toolbar = `rgb(${this.palette[3][0]},${this.palette[3][1]},${this.palette[3][2]})`;
              this.skin.dark = `rgb(${this.palette[4][0]},${this.palette[4][1]},${this.palette[4][2]})`;
              this.skin.wallpaper = `rgb(${this.palette[5][0]},${this.palette[5][1]},${this.palette[5][2]})`;
              this.skin.text = `rgb(${this.palette[6][0]},${this.palette[6][1]},${this.palette[6][2]})`;
              this.skin.panel = `rgb(${this.palette[7][0]},${this.palette[7][1]},${this.palette[7][2]})`;
              this.skin.button = `rgb(${this.palette[8][0]},${this.palette[8][1]},${this.palette[8][2]})`;
              this.skin.transparent = `rgb(${this.palette[0][0]},${this.palette[0][1]},${this.palette[0][2]})`;
              this.takeRootCSS = `
              :root {
                --sub-background : ${this.skin.background};
                --sub-card : ${this.skin.card};
                --sub-wallpaper : ${this.skin.wallpaper};
                --sub-toolbar : ${this.skin.toolbar};
                --sub-text : ${this.skin.text};
                --sub-dark : ${this.skin.dark};
                --sub-button : ${this.skin.button};
                --sub-panel : ${this.skin.panel};
                --transparent :  ${this.skin.transparent};
              }
              `}
            this.loadContext = true
            this.setContextCSS(this.projectName, this.takeRootCSS);
            //console.log("PALETA CARGADA");
          }
          catch (error) {
            setTimeout(() => {
              this.loadAvatarProject();
            }, 80)
          }
          return true;
        }
      }
    }
  }
</script>