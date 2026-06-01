import { Link } from "react-router-dom";
import LanguageSelector from "@/components/LanguageSelector";
import { LICENSES } from "@/constants/licenses";
import { useEffect, useRef } from "react";

/** Page d'accueil publique – reprend le contenu de renderHomePage() */
export default function HomePage() {
  const sectionsRef = useRef<(HTMLElement | null)[]>([]);

  // Animation au scroll (IntersectionObserver)
  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            const el = entry.target as HTMLElement;
            el.style.opacity = "1";
            el.style.transform = "translateY(0)";
          }
        });
      },
      { threshold: 0.1 },
    );

    sectionsRef.current.forEach((s) => {
      if (s) observer.observe(s);
    });

    return () => observer.disconnect();
  }, []);

  return (
    <div className="min-h-screen bg-gradient-to-b from-slate-900 via-slate-800 to-slate-900">
      {/* Nav */}
      <nav className="bg-slate-900/80 backdrop-blur-sm border-b border-slate-700/50 sticky top-0 z-50">
        <div className="max-w-6xl mx-auto px-4">
          <div className="flex items-center justify-between h-16">
            <div className="flex items-center gap-3">
              <span className="text-3xl">🛩️</span>
              <span className="text-xl font-bold text-white">AeroPath</span>
            </div>
            <div className="flex items-center gap-3">
              <LanguageSelector />
              <Link
                to="/login"
                className="px-5 py-2 text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition-all shadow-lg"
              >
                Connexion
              </Link>
              <Link
                to="/login#register"
                className="px-5 py-2 text-sm font-medium text-slate-300 border border-slate-600 hover:border-blue-500 hover:text-blue-400 rounded-lg transition-all"
              >
                S'inscrire
              </Link>
            </div>
          </div>
        </div>
      </nav>

      {/* Hero */}
      <section
        ref={(el) => { sectionsRef.current[0] = el; }}
        className="relative overflow-hidden"
        style={{ opacity: 0, transform: "translateY(20px)", transition: "all 0.6s ease-out" } as React.CSSProperties}
      >
        <div className="absolute inset-0 bg-gradient-to-b from-blue-500/5 via-transparent to-transparent pointer-events-none" />
        <div className="max-w-6xl mx-auto px-4 py-20 md:py-32">
          <div className="text-center max-w-3xl mx-auto">
            <div className="text-7xl md:text-8xl mb-8 animate-float">🛩️</div>
            <h1 className="text-4xl md:text-6xl font-extrabold text-white mb-6 leading-tight">
              Votre formation aéronautique,{" "}
              <span className="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-cyan-400">
                partout avec vous
              </span>
            </h1>
            <p className="text-lg md:text-xl text-slate-400 mb-10 max-w-2xl mx-auto leading-relaxed">
              Préparez vos licences PPL, LAPL, CPL, ATPL et IR avec des questions interactives,
              des leçons détaillées et un suivi personnalisé. Même sans connexion.
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Link
                to="/login#register"
                className="px-8 py-4 bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-xl text-lg transition-all shadow-xl hover:scale-105"
              >
                🚀 Commencer gratuitement
              </Link>
              <Link
                to="/login"
                className="px-8 py-4 bg-slate-800 hover:bg-slate-700 text-slate-300 font-semibold rounded-xl text-lg border border-slate-700 transition-all hover:scale-105"
              >
                🔐 J'ai déjà un compte
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* Features */}
      <section
        ref={(el) => { sectionsRef.current[1] = el; }}
        className="py-16 md:py-24"
        style={{ opacity: 0, transform: "translateY(20px)", transition: "all 0.6s ease-out" } as React.CSSProperties}
      >
        <div className="max-w-6xl mx-auto px-4">
          <h2 className="text-3xl md:text-4xl font-bold text-white text-center mb-4">Pourquoi AeroPath ?</h2>
          <p className="text-slate-400 text-center mb-16 max-w-xl mx-auto">
            Une plateforme complète conçue par des pilotes, pour les pilotes
          </p>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[
              { icon: "📚", title: "Questions illimitées", desc: "Des milliers de questions couvrant toutes les licences et tous les sujets, avec des explications détaillées." },
              { icon: "📓", title: "Leçons interactives", desc: "Apprenez à votre rythme avec des leçons structurées, quiz intégrés et suivi de progression." },
              { icon: "📊", title: "Suivi personnalisé", desc: "Statistiques détaillées, recommandations intelligentes et répétition espacée." },
              { icon: "📱", title: "Mode hors-ligne", desc: "Continuez à réviser même en vol ou dans les zones sans réseau. Synchronisation automatique." },
              { icon: "🧠", title: "Recommandations IA", desc: "Notre moteur analyse vos résultats et vous suggère les sujets à réviser." },
              { icon: "🎓", title: "Toutes les licences", desc: "PPL, LAPL, CPL, ATPL, IR – préparez toutes vos certifications au même endroit." },
            ].map((f, i) => (
              <div key={i} className="bg-slate-800/50 backdrop-blur-sm rounded-2xl p-6 border border-slate-700/50 hover:border-blue-500/30 transition-all group">
                <div className="w-14 h-14 bg-blue-500/10 rounded-xl flex items-center justify-center text-2xl mb-4 group-hover:bg-blue-500/20 transition-all">
                  {f.icon}
                </div>
                <h3 className="text-xl font-bold text-white mb-2">{f.title}</h3>
                <p className="text-slate-400 leading-relaxed">{f.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Stats */}
      <section
        ref={(el) => { sectionsRef.current[2] = el; }}
        className="py-16 bg-slate-800/30"
        style={{ opacity: 0, transform: "translateY(20px)", transition: "all 0.6s ease-out" } as React.CSSProperties}
      >
        <div className="max-w-6xl mx-auto px-4">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8 text-center">
            {[
              { n: "+5000", label: "Questions" },
              { n: "+200", label: "Leçons" },
              { n: "6", label: "Licences" },
              { n: "100%", label: "Hors-ligne" },
            ].map((s, i) => (
              <div key={i}>
                <div className="text-4xl font-bold text-white mb-1">{s.n}</div>
                <div className="text-slate-400 text-sm">{s.label}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Licences */}
      <section
        ref={(el) => { sectionsRef.current[3] = el; }}
        className="py-16 md:py-24"
        style={{ opacity: 0, transform: "translateY(20px)", transition: "all 0.6s ease-out" } as React.CSSProperties}
      >
        <div className="max-w-6xl mx-auto px-4">
          <h2 className="text-3xl md:text-4xl font-bold text-white text-center mb-4">Licences disponibles</h2>
          <p className="text-slate-400 text-center mb-16 max-w-xl mx-auto">
            Du pilote privé au transport aérien, nous couvrons toutes les étapes
          </p>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            {[
              { icon: "✈️", title: "PPL / LAPL", desc: "Pilote Privé – La base de l'aviation", gradient: "from-blue-900/20" },
              { icon: "🛩️", title: "CPL", desc: "Pilote Professionnel – Faites de votre passion un métier", gradient: "from-purple-900/20" },
              { icon: "✈️", title: "ATPL / IR", desc: "Transport Aérien & Vol aux Instruments – Le plus haut niveau", gradient: "from-amber-900/20" },
            ].map((l, i) => (
              <div key={i} className={`bg-gradient-to-br ${l.gradient} to-slate-800 rounded-xl p-6 border border-slate-800/30`}>
                <div className="text-3xl mb-3">{l.icon}</div>
                <h3 className="text-lg font-bold text-white mb-1">{l.title}</h3>
                <p className="text-sm text-slate-400">{l.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA */}
      <section
        ref={(el) => { sectionsRef.current[4] = el; }}
        className="py-16 md:py-24"
        style={{ opacity: 0, transform: "translateY(20px)", transition: "all 0.6s ease-out" } as React.CSSProperties}
      >
        <div className="max-w-4xl mx-auto px-4 text-center">
          <div className="bg-gradient-to-br from-blue-600/10 to-cyan-600/10 rounded-3xl p-10 md:p-16 border border-blue-500/20">
            <div className="text-6xl mb-6">🛩️</div>
            <h2 className="text-3xl md:text-4xl font-bold text-white mb-4">Prêt à décoller ?</h2>
            <p className="text-lg text-slate-400 mb-8 max-w-lg mx-auto">
              Rejoignez des centaines de pilotes qui préparent leur licence avec AeroPath
            </p>
            <Link
              to="/login#register"
              className="px-10 py-4 bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-xl text-lg transition-all shadow-xl hover:scale-105"
            >
              🚀 Créer mon compte gratuit
            </Link>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-slate-800 py-8">
        <div className="max-w-6xl mx-auto px-4 text-center">
          <div className="flex items-center justify-center gap-2 mb-4">
            <span className="text-xl">🛩️</span>
            <span className="font-bold text-white">AeroPath</span>
          </div>
          <p className="text-slate-500 text-sm">
            Formation aéronautique pour pilotes – PPL, LAPL, CPL, ATPL, IR
          </p>
          <p className="text-slate-600 text-xs mt-2">
            © {new Date().getFullYear()} AeroPath. Tous droits réservés.
          </p>
        </div>
      </footer>
    </div>
  );
}