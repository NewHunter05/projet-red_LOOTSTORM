package main

import "image/color"

const (
	screenWidth  = 960
	screenHeight = 600
	windowWidth  = 1440
	windowHeight = 900
)

var theme = struct {
	background color.RGBA
	panel      color.RGBA
	panelLight color.RGBA
	gold       color.RGBA
	goldBright color.RGBA
	text       color.RGBA
	muted      color.RGBA
	ruby       color.RGBA
}{
	background: color.RGBA{R: 15, G: 13, B: 18, A: 255},
	panel:      color.RGBA{R: 29, G: 25, B: 31, A: 245},
	panelLight: color.RGBA{R: 48, G: 39, B: 39, A: 245},
	gold:       color.RGBA{R: 157, G: 119, B: 52, A: 255},
	goldBright: color.RGBA{R: 230, G: 190, B: 93, A: 255},
	text:       color.RGBA{R: 238, G: 225, B: 193, A: 255},
	muted:      color.RGBA{R: 164, G: 148, B: 125, A: 255},
	ruby:       color.RGBA{R: 132, G: 48, B: 48, A: 255},
}

type classChoice struct {
	name string
	hp   int
}

var classes = []classChoice{
	{"Guerrier", 150}, {"Mage", 80}, {"Archer", 110},
	{"Assassin", 100}, {"Chevalier", 120}, {"Samourai", 115},
	{"Clerc", 105}, {"Barbare", 180}, {"Invocateur", 90},
}

type predefinedHero struct {
	name  string
	class string
	hp    int
}

var predefinedHeroes = []predefinedHero{
	{"Himiko Toga", "Assassin", 100}, {"Link", "Chevalier", 120},
	{"Frieren", "Mage", 80}, {"Musashi", "Samourai", 115},
	{"Elizabeth", "Clerc", 105}, {"Guts", "Barbare", 180},
	{"Sung Jin Woo", "Invocateur", 90}, {"Colley", "Archer", 110},
}

type shopItem struct {
	name  string
	price int
}

var shopItems = []shopItem{
	{"Potion de vie", 3}, {"Potion de poison", 6}, {"Potion de mana", 6},
	{"Livre de Sort : Boule de Feu", 25}, {"Livre de Sort : Soin", 25},
	{"Livre de Sort : Régénération", 25}, {"Livre de Sort : Poison", 25},
	{"Livre de Sort : Brûlure", 25}, {"Livre de Sort : Éclair", 25},
	{"Livre de Sort : Barrière Sacrée", 25}, {"Fourrure de Loup", 4},
	{"Peau de Troll", 7}, {"Cuir de Sanglier", 3}, {"Plume de Corbeau", 1},
	{"Amelioration d'inventaire", 30},
}

var forgeItems = []string{
	"Chapeau de l'aventurier", "Tunique de l'aventurier", "Bottes de l'aventurier",
}

var spriteFiles = map[string]string{
	"class:Guerrier": "class_guerrier.png", "class:Mage": "class_mage.png", "class:Archer": "class_archer.png",
	"class:Assassin": "class_assassin.png", "class:Chevalier": "class_chevalier.png", "class:Samourai": "class_samourai.png",
	"class:Clerc": "class_clerc.png", "class:Barbare": "class_barbare.png", "class:Invocateur": "class_invocateur.png",
	"hero:Himiko Toga": "hero_himiko_toga.png", "hero:Link": "hero_link.png",
	"hero:Frieren": "hero_frieren.png", "hero:Musashi": "hero_musashi.png",
	"hero:Elizabeth": "hero_elizabeth.png", "hero:Guts": "hero_guts.png", "hero:Sung Jin Woo": "hero_sung_jin_woo.png",
	"hero:Colley": "hero_colley.png",
	"goblin":      "enemy_goblin.png", "slime": "enemy_slime.png", "ghost": "enemy_ghost.png",
	"golem": "enemy_golem.png", "troll": "enemy_troll.png", "skeleton": "enemy_skeleton.png",
	"wolf": "enemy_wolf.png", "wizard": "enemy_wizard.png", "duck": "enemy_duck.png",
	"dragon": "enemy_dragon.png", "summon": "enemy_summon.png",
}
