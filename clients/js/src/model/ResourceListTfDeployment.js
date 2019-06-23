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
    define(['ApiClient', 'model/ResourceTfDeployment', 'model/ResourceTfStack'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./ResourceTfDeployment'), require('./ResourceTfStack'));
  } else {
    // Browser globals (root is window)
    if (!root.Lunaform) {
      root.Lunaform = {};
    }
    root.Lunaform.ResourceListTfDeployment = factory(root.Lunaform.ApiClient, root.Lunaform.ResourceTfDeployment, root.Lunaform.ResourceTfStack);
  }
}(this, function(ApiClient, ResourceTfDeployment, ResourceTfStack) {
  'use strict';




  /**
   * The ResourceListTfDeployment model module.
   * @module model/ResourceListTfDeployment
   * @version 0.0.1-alpha
   */

  /**
   * Constructs a new <code>ResourceListTfDeployment</code>.
   * @alias module:model/ResourceListTfDeployment
   * @class
   * @param deployments {Array.<module:model/ResourceTfDeployment>} 
   * @param stack {module:model/ResourceTfStack} 
   */
  var exports = function(deployments, stack) {
    var _this = this;

    _this['deployments'] = deployments;
    _this['stack'] = stack;
  };

  /**
   * Constructs a <code>ResourceListTfDeployment</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ResourceListTfDeployment} obj Optional instance to populate.
   * @return {module:model/ResourceListTfDeployment} The populated <code>ResourceListTfDeployment</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('deployments')) {
        obj['deployments'] = ApiClient.convertToType(data['deployments'], [ResourceTfDeployment]);
      }
      if (data.hasOwnProperty('stack')) {
        obj['stack'] = ResourceTfStack.constructFromObject(data['stack']);
      }
    }
    return obj;
  }

  /**
   * @member {Array.<module:model/ResourceTfDeployment>} deployments
   */
  exports.prototype['deployments'] = undefined;
  /**
   * @member {module:model/ResourceTfStack} stack
   */
  exports.prototype['stack'] = undefined;



  return exports;
}));


