"use strict";$R.add("plugin","alignment",{translations:{en:{align:"Align","align-left":"Align Left","align-center":"Align Center","align-right":"Align Right","align-justify":"Align Justify"}},init:function(t){this.app=t,this.opts=t.opts,this.lang=t.lang,this.block=t.block,this.selection=t.selection,this.toolbar=t.toolbar},start:function(){var t={};t.left={title:this.lang.get("align-left"),api:"plugin.alignment.set",args:"left"},t.center={title:this.lang.get("align-center"),api:"plugin.alignment.set",args:"center"},t.right={title:this.lang.get("align-right"),api:"plugin.alignment.set",args:"right"},t.justify={title:this.lang.get("align-justify"),api:"plugin.alignment.set",args:"justify"};var i=this.toolbar.addButton("alignment",{title:this.lang.get("align")});i.setIcon('<i class="re-icon-alignment"></i>'),i.setDropdown(t)},set:function(i){this.selection.getBlocks().forEach(function(t){$R.dom(t).removeClass("rd-text-left rd-text-right rd-text-center rd-text-justify").addClass("rd-text-"+i)})}});