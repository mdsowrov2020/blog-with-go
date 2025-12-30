package repo

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type PostRepo interface {
	Create(p Post) (*Post, error)
	List() ([]*Post, error)
	Get(id int) (*Post, error)
	Update(p Post) (*Post, error)
	Delete(id int) error
}

type postRepo struct {
	postList []*Post
}

func NewPostRepo() PostRepo {
	repo := &postRepo{}
	generateInitialPosts(repo)
	return repo
}

func (r *postRepo) Create(p Post) (*Post, error) {
	p.ID = len(r.postList) + 1
	r.postList = append(r.postList, &p)
	return &p, nil
}

func (r *postRepo) List() ([]*Post, error) {
	return r.postList, nil
}

func (r *postRepo) Get(id int) (*Post, error) {
	for _, post := range r.postList {
		if post.ID == id {
			return post, nil
		}
	}

	return nil, nil
}

func (r *postRepo) Update(p Post) (*Post, error) {
	for idx, product := range r.postList {
		if product.ID == p.ID {
			r.postList[idx] = &p
		}
	}

	return &p, nil
}

func (r *postRepo) Delete(id int) error {
	var tempList []*Post

	for _, post := range r.postList {
		if post.ID != id {
			tempList = append(tempList, post)
		}
	}

	r.postList = tempList

	return nil
}

func generateInitialPosts(r *postRepo) {
	post1 := &Post{
		ID:          1,
		Title:       "Mobile-First CSS: Is It Time for a Rethink?",
		Description: "The mobile-first design methodology is great—it focuses on what really matters to the user, it’s well-practiced, and it’s been a common design pattern for years. So developing your CSS mobile-first should also be great, too…right? ",
		ImageURL:    "https://lh4.googleusercontent.com/O8lxNeIY3Hb0YDs2EP7QFhGdGsBXOG7mSTCdAJBd5xkm-6RwrpkS1BN63W7RurVCP3nOH9sNpAR9JNGvIGnUTzG0NYm4sUqI5bU2QPhXYEawmKfeUJ_6YwWAIid2ZDHEdRzaQ1LxzUNTGbGk5g",
	}
	post2 := &Post{
		ID:          2,
		Title:       "I am a creative.",
		Description: "I am a creative. What I do is alchemy. It is a mystery. I do not so much do it, as let it be done through me. ",
		ImageURL:    "",
	}

	r.postList = append(r.postList, post1)
	r.postList = append(r.postList, post2)
}
