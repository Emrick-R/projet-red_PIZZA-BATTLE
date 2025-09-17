🍕 Pizza Quest RPG
Un jeu de rôle thématique pizza développé en Go

📋 Description
Pizza Quest RPG est un jeu de rôle en ligne de commande où vous incarnez un aventurier culinaire dans un monde peuplé de chefs italiens. Collectez des ingrédients, préparez des potions gourmandes et devenez le maître pizza ultime !

🚀 Installation et Lancement
Cloner le projet :
git clone https://github.com/Emrick-R/projet-red_PIZZA-BATTLE.git

# Lancer le jeu
go run main.go
🎮 Fonctionnalités du Jeu
⚔️ Système de Combat

Ennemis : Chefs italiens de différents niveaux
Types d'ennemis :

🟢 Facile : Petit Giovanni
🟡 Normal : El Don Pastabox 3000
🔴 Boss : Ultra Mega Hyper Giovanni EX Turbo GX



🏃‍♂️ Statistiques du Personnage



🎒 Système d'Inventaire

Frigo portatif : Inventaire limité en taille
Objets disponibles :

🧪 Séance de Sport : Potion de vie
🧪 Tabasco : Potion de Poison
📚 Livres de sorts et recettes :
🍖 Ingrédients (Peau de Loup, Cuir de Sanglier, etc.) :
👨‍🍳 Équipements de chef :



🏪 Commerce

Marchand : Super Marché : Achat/vente d'objets
Oncle Pepito : Forgeron Craft d'objets

⚡ Compétences

Coup de poing : 10 dégâts - 0 Mana
Boule de feu : 20 dégâts (sort magique)
Sort de la Mort Qui Tue : One Shot l'ennemi

🎯 Équipements

👒 Chapeau de l'Aventurier : Toque de Chef
👔 Tablier
🥾 Bottes d'Aventurier 
🔪 Couteaux de cuisine
🍕 Rouleau à pizza

🛠️ Structure du Code
Phase 1 - Fonctionnalités Implémentées ✅

Structure Character - Définition du personnage
initCharacter() - Initialisation dans main()
displayInfo() - Affichage des informations
accessInventory() - Accès à l'inventaire
takePot() - Utilisation des potions
Menu Principal - Interface switch/case
Gestion Inventaire :

addInventory() - Ajouter objets
removeInventory() - Retirer objets


isDead() - Vérification état de vie
poisonPot() - Système de poison temporel
Système de Skills :

Skills de base
Skill books
spellbook() - Utilisation des sorts


characterCreation() - Création personnalisée
Limite d'inventaire - Gestion du poids

🎨 Fonctionnalités Spéciales
🍄 Système de Poison

Tabasco : Potion de poison thématique
Effet temporel avec bibliothèque time
Dégâts progressifs 10 dégâts par tour à l'ennemi

🎲 Combat : à modifier

Options :
Easter egg découvert

📊 Système de Progression

Gain d'expérience par combat
Montée en niveau automatique
Amélioration des statistiques

🔧 Commandes de Développement
go run main.go

Ne pas atteindre un certain poids kilogramme sinon la mort vous attend

📝 Notes de Développement

Interface utilisateur avec emojis intégrés
Menus structurés selon le package d'affichage
Commentaires détaillés dans le code
Architecture modulaire et extensible

Bon appétit et bonne aventure dans pizza battle ! 🍕⚔️