/* eslint-env browser */
'use strict'

var $ = require('jquery')
const Web3 = require('web3')

require('bootstrap/dist/css/bootstrap.min.css')

const App = {
  web3: null,
  contracts: {},

  init: async function () {
    // Load pets.
    $.getJSON('pets.json', function (data) {
      var petsRow = $('#petsRow')
      var petTemplate = $('#petTemplate')

      for (let i = 0; i < data.length; i++) {
        petTemplate.find('.panel-title').text(data[i].name)
        petTemplate.find('img').attr('src', data[i].picture)
        petTemplate.find('.pet-breed').text(data[i].breed)
        petTemplate.find('.pet-age').text(data[i].age)
        petTemplate.find('.pet-location').text(data[i].location)
        petTemplate.find('.btn-adopt').attr('data-id', data[i].id)

        petsRow.append(petTemplate.html())
      }
    })

    return await App.initWeb3()
  },

  initWeb3: async function () {
    var standardProvider = null
    if (window.ethereum) {
      standardProvider = window.ethereum
      try {
        // Request account access if needed
        await window.ethereum.enable()
      } catch (error) {
        console.log('User denied account access')
      }
    } else if (window.web3) {
      // Legacy dapp browsers...
      standardProvider = window.web3.currentProvider
    } else {
      // Non-dapp browsers...
      console.log(
        'Non-Ethereum browser detected. You should consider trying MetaMask!'
      )
    }

    App.web3 = new Web3(standardProvider) // eslint-disable-line require-atomic-updates

    return App.initContract()
  },

  initContract: async function () {
    const adoption = require('../build/contracts/Adoption.json')

    const netid = await App.web3.eth.net.getId()

    App.contracts.Adoption = new App.web3.eth.Contract(
      adoption.abi,
      adoption.networks[netid].address
    )

    // Use our contract to retrieve and mark the adopted pets
    App.markAdopted()

    return App.bindEvents()
  },

  bindEvents: function () {
    $(document).on('click', '.btn-adopt', App.handleAdopt)
  },

  markAdopted: async function () {
    try {
      const adopters = await App.contracts.Adoption.methods.getAdopters().call()
      for (let i = 0; i < adopters.length; i++) {
        if (adopters[i] !== '0x0000000000000000000000000000000000000000') {
          $('.panel-pet')
            .eq(i)
            .find('button')
            .text('Success')
            .attr('disabled', true)
        }
      }
    } catch (err) {
      console.log(err.message)
    }
  },

  handleAdopt: async function (event) {
    event.preventDefault()
    var petId = parseInt($(event.target).data('id'))

    try {
      const accounts = await App.web3.eth.getAccounts()
      await App.contracts.Adoption.methods
        .adopt(petId)
        .send({ from: accounts[0] })
      await App.markAdopted()
    } catch (err) {
      console.log('Error calling adopt', err.message)
    }
  },
}

window.addEventListener('DOMContentLoaded', App.init)
