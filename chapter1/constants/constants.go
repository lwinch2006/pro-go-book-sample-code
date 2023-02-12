package constants

// Shared templates
const LayoutFolder string = "views/shared/"
const MainLayout string = LayoutFolder + "layout"

// Home pages
const HomeViewsFolder string = "views/home/"
const HomeFormPage = HomeViewsFolder + "form"
const HomeListPage = HomeViewsFolder + "list"
const HomeSorryPage = HomeViewsFolder + "sorry"
const HomeThanksPage = HomeViewsFolder + "thanks"
const HomeWelcomePage = HomeViewsFolder + "welcome"

func GetHomePages() []string {
	return []string{HomeFormPage, HomeListPage, HomeSorryPage, HomeThanksPage, HomeWelcomePage}
}
