package desen

import "github.com/fusugongzi/gotools/slices"

var LTPrefix = []string{"130", "131", "132", "145", "155", "156", "176", "185", "186"}
var YDPrefix = []string{"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "178", "182", "183", "184", "187", "188", "198"}
var DXPrefix = []string{"133", "153", "177", "180", "181", "189", "199"}
var VRPrefix = []string{"162", "165", "167", "170", "171", "184"}
var AllMobilePrefix = slices.Merge(LTPrefix, YDPrefix, DXPrefix, VRPrefix)

var Digits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
