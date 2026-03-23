import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <main className="flex-1">
        <HeroSection />
        <FeaturesSection />
      </main>
      <Footer />
    </div>
  );
}

function Header() {
  return (
    <header className="border-b bg-background/80 backdrop-blur-sm sticky top-0 z-50">
      <div className="max-w-6xl mx-auto px-6 h-16 flex items-center justify-between">
        <span className="text-xl font-bold tracking-tight">Menta Clone</span>
        <nav className="hidden sm:flex items-center gap-6 text-sm text-muted-foreground">
          <a href="#" className="hover:text-foreground transition-colors">
            メンターを探す
          </a>
          <a href="#" className="hover:text-foreground transition-colors">
            プランを見る
          </a>
        </nav>
        <Button size="sm">無料で始める</Button>
      </div>
    </header>
  );
}

function HeroSection() {
  return (
    <section className="py-24 px-6 text-center bg-gradient-to-b from-muted/40 to-background">
      <div className="max-w-3xl mx-auto flex flex-col items-center gap-6">
        <span className="text-sm font-medium px-3 py-1 rounded-full bg-primary/10 text-primary">
          メンタリングプラットフォーム
        </span>
        <h1 className="text-5xl font-bold tracking-tight leading-tight">
          あなたの成長を
          <br />
          加速するメンターを見つけよう
        </h1>
        <p className="text-lg text-muted-foreground max-w-xl">
          エンジニア・デザイナー・ビジネスパーソンなど、各分野のプロフェッショナルがあなたのキャリアをサポートします。
        </p>
        <div className="flex flex-col sm:flex-row gap-3 mt-2">
          <Button size="lg">メンターを探す</Button>
          <Button size="lg" variant="outline">
            メンターとして登録する
          </Button>
        </div>
      </div>
    </section>
  );
}

const features = [
  {
    title: "豊富なメンター陣",
    description:
      "エンジニアリング・デザイン・マーケティングなど、幅広い分野のプロフェッショナルが在籍しています。",
  },
  {
    title: "柔軟なプラン",
    description:
      "単発のレビューから継続的なメンタリングまで、あなたのペースや目標に合わせたプランを選べます。",
  },
  {
    title: "安心のサポート",
    description:
      "実績・レビューの確認や、専任サポートチームへの相談など、安心して利用できる仕組みを整えています。",
  },
];

function FeaturesSection() {
  return (
    <section className="py-20 px-6">
      <div className="max-w-6xl mx-auto">
        <h2 className="text-3xl font-bold text-center mb-12">選ばれる理由</h2>
        <div className="grid sm:grid-cols-3 gap-8">
          {features.map((feature) => (
            <div
              key={feature.title}
              className="rounded-xl border bg-card p-6 flex flex-col gap-3"
            >
              <h3 className="text-lg font-semibold">{feature.title}</h3>
              <p className="text-sm text-muted-foreground leading-relaxed">
                {feature.description}
              </p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

function Footer() {
  return (
    <footer className="border-t py-8 px-6 text-center text-sm text-muted-foreground">
      © 2026 Menta Clone. All rights reserved.
    </footer>
  );
}
