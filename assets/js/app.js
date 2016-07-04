new Vue({
  el: '#orders',
  data: {
    order: {
      type: 'order',
      attributes: {
        product_title: '',
        product_description: '',
      }
    },
    orders: [],
    },
  ready: function() {
    this.orderEndpoint = this.$resource('api/v1/order/{id}');
    this.orderDeleteEndpoint = this.$resource('api/v1/order/{id}/delete');
    this.orderPatchEndpoint = this.$resource('api/v1/order/{id}/patch');
    this.fetchOrders();
  },
  methods: {
    fetchOrders: function() {
      this.orderEndpoint.get().then(function(response) {
        this.$set('orders', response.data.data);
      }, function(response) {
        console.log(response);
      });
    },
    postOrder: function() {
        this.orderEndpoint.save({data: this.order}).then(function(response) {
            this.orders.push(response.data.data);
        }, function(response) {
            console.log(response);
        });

        this.order.attributes = {
            product_title: '',
            product_description: '',
        };
    },
    deleteOrder: function(index) {
        this.orderDeleteEndpoint.get({id: this.orders[index].attributes.id}).then(function(response) {
            this.orders.splice(index, 1);
        }, function(response) {
            console.log(response);
        });
    }
  }
});
