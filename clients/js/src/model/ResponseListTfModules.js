/**
 * lunaform
 * This is a RESTful server for managing Terraform plan and apply jobs and the auditing of actions to approve those apply jobs. The inspiration for this project is the AWS CloudFormation API's. The intention is to implement a locking mechanism not only for the terraform state, but for the plan and apply of terraform modules. Once a `module` plan starts, it is instantiated as a `stack` within the nomencalture of `lunaform`. 
 *
 * OpenAPI spec version: 0.0.1-alpha
 * Contact: drew.sonne@gmail.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/HalRscLinks', 'model/ResourceListTfModule'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./HalRscLinks'), require('./ResourceListTfModule'));
  } else {
    // Browser globals (root is window)
    if (!root.Lunaform) {
      root.Lunaform = {};
    }
    root.Lunaform.ResponseListTfModules = factory(root.Lunaform.ApiClient, root.Lunaform.HalRscLinks, root.Lunaform.ResourceListTfModule);
  }
}(this, function(ApiClient, HalRscLinks, ResourceListTfModule) {
  'use strict';




  /**
   * The ResponseListTfModules model module.
   * @module model/ResponseListTfModules
   * @version 0.0.1-alpha
   */

  /**
   * Constructs a new <code>ResponseListTfModules</code>.
   * List of tf modules
   * @alias module:model/ResponseListTfModules
   * @class
   * @param links {module:model/HalRscLinks} 
   * @param embedded {module:model/ResourceListTfModule} 
   */
  var exports = function(links, embedded) {
    var _this = this;

    _this['_links'] = links;
    _this['_embedded'] = embedded;
  };

  /**
   * Constructs a <code>ResponseListTfModules</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ResponseListTfModules} obj Optional instance to populate.
   * @return {module:model/ResponseListTfModules} The populated <code>ResponseListTfModules</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('_links')) {
        obj['_links'] = HalRscLinks.constructFromObject(data['_links']);
      }
      if (data.hasOwnProperty('_embedded')) {
        obj['_embedded'] = ResourceListTfModule.constructFromObject(data['_embedded']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/HalRscLinks} _links
   */
  exports.prototype['_links'] = undefined;
  /**
   * @member {module:model/ResourceListTfModule} _embedded
   */
  exports.prototype['_embedded'] = undefined;



  return exports;
}));


