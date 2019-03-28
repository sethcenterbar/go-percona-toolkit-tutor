package structs

type Toolbox struct {
	Tools []Tool `yaml:"tools"`
}

type Tool struct {
	Name      string     `yaml:"name"`
	Summary   string     `yaml:"summary"`
	Situation string     `yaml:"situation"`
	Examples  []Example  `yaml:"examples"`
	BlogPosts []BlogPost `yaml:"blog_posts"`
	Videos    []Video    `yaml:"videos"`
}

type Example struct {
	Description string `yaml:"description"`
	Command     string `yaml:"command"`
}

type BlogPost struct {
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
}

type Video struct {
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
}
