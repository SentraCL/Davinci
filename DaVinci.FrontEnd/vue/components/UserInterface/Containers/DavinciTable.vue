<template>
  <span>
    <div :style="minHeight" class="row">
      <table class="table">
        <thead>
          <tr>
            <th v-for="key in headers" @click="sortBy(key)" :key="key" :class="{ active: sortKey == key }">
              {{ key | capitalize }}
              <span class="arrow" :class="sortOrders[key] > 0 ? 'asc' : 'dsc'">
              </span>
            </th>
          </tr>
        </thead>
        <tbody style="height:150px">
          <tr v-for="entry in filteredData" :key="entry.key">
            <td v-for="key in headers" :key="key" v-html="entry[key]" class="td">

            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <hr />
    <div class="row" v-if="!Inactive">
      <uib-pagination v-if="isPagination" @change="changePage()" firstText="«" lastText="»" previousText="‹" nextText="›" :total-items="data.length" v-model="toPage" :itemsPerPage="pagination.itemsPerPage" :max-size="15" :boundary-links="true">
      </uib-pagination>
    </div>

  </span>
</template>

<script>
  import Vue from "vue";
  var uibPagination = require("vuejs-uib-pagination");
  Vue.use(uibPagination);

  export default {
    props: {
      data: Array,
      columns: Array,
      filterKey: String,
      pagination: {},
      Inactive:false
    },
    data: function () {
      console.log("data",this.data)
      console.log("columns",this.columns)
      console.log("filterKey",this.filterKey)
      var sortOrders = {};
      var _headers = [];
      if (this.columns == null) {
        _headers = Object.keys(this.data[0]);
      } else {
        _headers = this.columns;
      }
      _headers.forEach(function (key) {
        sortOrders[key] = 1;
      });
      return {
        sortKey: "",
        sortOrders: sortOrders,
        headers: _headers,
        isPagination: this.pagination != null && this.pagination.itemsPerPage <= this.data.length,
        toPage: this.pagination
      };
    },
    watch: {
      data(newValue, oldValue) {
        var sortOrders = {};
        var _headers = [];
        if (this.columns == null) {
          _headers = Object.keys(this.data[0]);
        } else {
          _headers = this.columns;
        }
        _headers.forEach(function (key) {
          sortOrders[key] = 1;
        });

        this.sortKey = "";
        this.sortOrders = sortOrders;
        this.headers = _headers;
        this.isPagination = this.pagination != null && this.pagination.itemsPerPage <= this.data.length;
        this.toPage = this.pagination;
      }
    },
    computed: {
      minHeight: function () {
        if (this.isPagination) {
          return `min-height:${parseInt(this.pagination.itemsPerPage, 10) * 65}px !important;height:${parseInt(this.pagination.itemsPerPage, 10) * 65}px !important;`;
        } else {
          return "";
        }
      },
      filteredData: function () {
        var sortKey = this.sortKey;
        var filterKey = this.filterKey && this.filterKey.toLowerCase();
        var order = this.sortOrders[sortKey] || 1;
        var data = this.data;
        if (filterKey) {
          data = data.filter(function (row) {
            return Object.keys(row).some(function (key) {
              return (
                String(row[key])
                  .toLowerCase()
                  .indexOf(filterKey) > -1
              );
            });
          });
        }
        if (sortKey) {
          data = data.slice().sort(function (a, b) {
            a = a[sortKey];
            b = b[sortKey];
            return (a === b ? 0 : a > b ? 1 : -1) * order;
          });
        }

        if (this.isPagination) {
          if (Object.keys(this.toPage).indexOf("currentPage") == -1) {
            this.toPage.currentPage = 1;
          }

          var dataPag = []
          var ixp = this.pagination.itemsPerPage;

          var end = data.length;
          var begin = (end <= ixp) ? 0 : ((this.toPage.currentPage * ixp) - ixp);
          var index = 0;

          for (var key in data) {
            var item = data[key];
            if (index >= begin) {
              dataPag.push(item);
              if (dataPag.length == ixp + 1) {
                break;
              }
            }
            index++
          }
          return dataPag;
        } else {
          return data;
        }
      }
    },
    filters: {
      capitalize: function (str) {
        return str.charAt(0).toUpperCase() + str.slice(1);
      }
    },
    methods: {
      changePage() {
        this.$emit('changePage')
      },
      sortBy: function (key) {
        this.sortKey = key;
        this.sortOrders[key] = this.sortOrders[key] * -1;
      }
    }
  };
</script>
<style scoped>
  .td{
    min-width: 60px;
  }

  table{
    font-size: 14px;
  }
</style>